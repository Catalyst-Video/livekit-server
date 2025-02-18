package rtc

import (
	"errors"
	"sync"
	"time"

	"github.com/livekit/protocol/utils"
	"github.com/pion/ion-sfu/pkg/buffer"
	"github.com/pion/ion-sfu/pkg/sfu"
	"github.com/pion/ion-sfu/pkg/twcc"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/rtcerr"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/logger"
	"github.com/livekit/livekit-server/pkg/rtc/types"
	livekit "github.com/livekit/livekit-server/proto"
)

var (
	feedbackTypes = []webrtc.RTCPFeedback{
		{Type: webrtc.TypeRTCPFBGoogREMB},
		{Type: webrtc.TypeRTCPFBNACK},
		{Type: webrtc.TypeRTCPFBNACK, Parameter: "pli"}}
)

// MediaTrack represents a WebRTC track that needs to be forwarded
// Implements the PublishedTrack interface
type MediaTrack struct {
	params      MediaTrackParams
	ssrc        webrtc.SSRC
	name        string
	streamID    string
	kind        livekit.TrackType
	codec       webrtc.RTPCodecParameters
	muted       utils.AtomicFlag
	simulcasted bool

	// channel to send RTCP packets to the source
	lock sync.RWMutex
	// map of target participantId -> *SubscribedTrack
	subscribedTracks map[string]*SubscribedTrack
	twcc             *twcc.Responder
	audioLevel       *AudioLevel
	receiver         sfu.Receiver
	lastPLI          time.Time

	onClose func()
}

type MediaTrackParams struct {
	TrackID        string
	ParticipantID  string
	RTCPChan       chan []rtcp.Packet
	BufferFactory  *buffer.Factory
	ReceiverConfig ReceiverConfig
	AudioConfig    config.AudioConfig
	Stats          *RoomStatsReporter
	Width          uint32
	Height         uint32
}

func NewMediaTrack(track *webrtc.TrackRemote, params MediaTrackParams) *MediaTrack {
	t := &MediaTrack{
		params:           params,
		ssrc:             track.SSRC(),
		streamID:         track.StreamID(),
		kind:             ToProtoTrackKind(track.Kind()),
		codec:            track.Codec(),
		subscribedTracks: make(map[string]*SubscribedTrack),
	}

	return t
}

func (t *MediaTrack) Start() {
}

func (t *MediaTrack) ID() string {
	return t.params.TrackID
}

func (t *MediaTrack) Kind() livekit.TrackType {
	return t.kind
}

func (t *MediaTrack) Name() string {
	return t.name
}

func (t *MediaTrack) IsMuted() bool {
	return t.muted.Get()
}

func (t *MediaTrack) SetMuted(muted bool) {
	t.muted.TrySet(muted)

	// mute all of the subscribedtracks
	t.lock.RLock()
	for _, st := range t.subscribedTracks {
		st.SetPublisherMuted(muted)
	}
	t.lock.RUnlock()
}

func (t *MediaTrack) OnClose(f func()) {
	t.onClose = f
}

func (t *MediaTrack) IsSubscriber(subId string) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.subscribedTracks[subId] != nil
}

// AddSubscriber subscribes sub to current mediaTrack
func (t *MediaTrack) AddSubscriber(sub types.Participant) error {
	if !sub.CanSubscribe() {
		return ErrPermissionDenied
	}

	t.lock.Lock()
	defer t.lock.Unlock()
	existingSt := t.subscribedTracks[sub.ID()]

	// don't subscribe to the same track multiple times
	if existingSt != nil {
		return nil
	}

	if t.receiver == nil {
		// cannot add, no receiver
		return errors.New("cannot subscribe without a receiver in place")
	}

	codec := t.receiver.Codec()
	if err := sub.SubscriberMediaEngine().RegisterCodec(codec, t.receiver.Kind()); err != nil {
		return err
	}

	// using DownTrack from ion-sfu
	streamId := t.params.ParticipantID
	if sub.ProtocolVersion().SupportsPackedStreamId() {
		// when possible, pack both IDs in streamID to allow new streams to be generated
		// react-native-webrtc still uses stream based APIs and require this
		streamId = PackStreamID(t.params.ParticipantID, t.ID())
	}
	receiver := NewWrappedReceiver(t.receiver, t.ID(), streamId)
	downTrack, err := sfu.NewDownTrack(webrtc.RTPCodecCapability{
		MimeType:     codec.MimeType,
		ClockRate:    codec.ClockRate,
		Channels:     codec.Channels,
		SDPFmtpLine:  codec.SDPFmtpLine,
		RTCPFeedback: feedbackTypes,
	}, receiver, t.params.BufferFactory, sub.ID(), t.params.ReceiverConfig.packetBufferSize)
	if err != nil {
		return err
	}
	subTrack := NewSubscribedTrack(downTrack)

	transceiver, err := sub.SubscriberPC().AddTransceiverFromTrack(downTrack, webrtc.RTPTransceiverInit{
		Direction: webrtc.RTPTransceiverDirectionSendonly,
	})
	if err != nil {
		return err
	}

	downTrack.SetTransceiver(transceiver)
	// when outtrack is bound, start loop to send reports
	downTrack.OnBind(func() {
		subTrack.SetPublisherMuted(t.IsMuted())
		go t.sendDownTrackBindingReports(sub)
	})

	downTrack.OnCloseHandler(func() {
		go func() {
			t.lock.Lock()
			delete(t.subscribedTracks, sub.ID())
			t.lock.Unlock()

			t.params.Stats.SubSubscribedTrack(t.kind.String())

			// ignore if the subscribing sub is not connected
			if sub.SubscriberPC().ConnectionState() == webrtc.PeerConnectionStateClosed {
				return
			}

			// if the source has been terminated, we'll need to terminate all of the subscribedtracks
			// however, if the dest sub has disconnected, then we can skip
			sender := transceiver.Sender()
			if sender == nil {
				return
			}
			logger.Debugw("removing peerconnection track",
				"track", t.params.TrackID,
				"participantId", t.params.ParticipantID,
				"destParticipant", sub.Identity())
			if err := sub.SubscriberPC().RemoveTrack(sender); err != nil {
				if err == webrtc.ErrConnectionClosed {
					// sub closing, can skip removing subscribedtracks
					return
				}
				if _, ok := err.(*rtcerr.InvalidStateError); !ok {
					logger.Warnw("could not remove remoteTrack from forwarder", err,
						"sub", sub.Identity())
				}
			}

			sub.RemoveSubscribedTrack(t.params.ParticipantID, subTrack)
			sub.Negotiate()
		}()
	})

	t.subscribedTracks[sub.ID()] = subTrack

	t.receiver.AddDownTrack(downTrack, t.shouldStartWithBestQuality())
	// since sub will lock, run it in a gorountine to avoid deadlocks
	go func() {
		sub.AddSubscribedTrack(t.params.ParticipantID, subTrack)
		sub.Negotiate()
	}()

	t.params.Stats.AddSubscribedTrack(t.kind.String())
	return nil
}

// AddReceiver adds a new RTP receiver to the track
func (t *MediaTrack) AddReceiver(receiver *webrtc.RTPReceiver, track *webrtc.TrackRemote, twcc *twcc.Responder) {
	t.lock.Lock()
	defer t.lock.Unlock()

	buff, rtcpReader := t.params.BufferFactory.GetBufferPair(uint32(track.SSRC()))
	buff.OnFeedback(func(fb []rtcp.Packet) {
		if t.params.Stats != nil {
			t.params.Stats.incoming.HandleRTCP(fb)
		}
		// feedback for the source RTCP
		t.params.RTCPChan <- fb
	})

	if t.Kind() == livekit.TrackType_AUDIO {
		t.audioLevel = NewAudioLevel(t.params.AudioConfig.ActiveLevel, t.params.AudioConfig.MinPercentile)
		buff.OnAudioLevel(func(level uint8) {
			t.audioLevel.Observe(level)
		})
	} else if t.Kind() == livekit.TrackType_VIDEO {
		if twcc != nil {
			buff.OnTransportWideCC(func(sn uint16, timeNS int64, marker bool) {
				twcc.Push(sn, timeNS, marker)
			})
		}
	}

	rtcpReader.OnPacket(func(bytes []byte) {
		pkts, err := rtcp.Unmarshal(bytes)
		if err != nil {
			logger.Errorw("could not unmarshal RTCP", err)
			return
		}

		for _, pkt := range pkts {
			switch pkt := pkt.(type) {
			case *rtcp.SourceDescription:
			// do nothing for now
			case *rtcp.SenderReport:
				buff.SetSenderReportData(pkt.RTPTime, pkt.NTPTime)
			}
		}
	})

	if t.receiver == nil {
		t.receiver = sfu.NewWebRTCReceiver(receiver, track, t.params.ParticipantID, sfu.WithPliThrottle(0))
		t.receiver.SetRTCPCh(t.params.RTCPChan)
		t.receiver.OnCloseHandler(func() {
			t.lock.Lock()
			t.receiver = nil
			onclose := t.onClose
			t.lock.Unlock()
			t.RemoveAllSubscribers()
			t.params.Stats.SubPublishedTrack(t.kind.String())
			if onclose != nil {
				onclose()
			}
		})
		t.params.Stats.AddPublishedTrack(t.kind.String())
	}
	t.receiver.AddUpTrack(track, buff, t.shouldStartWithBestQuality())
	// when RID is set, track is simulcasted
	t.simulcasted = track.RID() != ""

	buff.Bind(receiver.GetParameters(), buffer.Options{
		MaxBitRate: t.params.ReceiverConfig.maxBitrate,
	})
}

// RemoveSubscriber removes participant from subscription
// stop all forwarders to the client
func (t *MediaTrack) RemoveSubscriber(participantId string) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if subTrack := t.subscribedTracks[participantId]; subTrack != nil {
		go subTrack.DownTrack().Close()
	}
}

func (t *MediaTrack) RemoveAllSubscribers() {
	logger.Debugw("removing all subscribers", "track", t.params.TrackID)
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, subTrack := range t.subscribedTracks {
		go subTrack.DownTrack().Close()
	}
	t.subscribedTracks = make(map[string]*SubscribedTrack)
}

func (t *MediaTrack) ToProto() *livekit.TrackInfo {
	return &livekit.TrackInfo{
		Sid:       t.ID(),
		Type:      t.Kind(),
		Name:      t.Name(),
		Muted:     t.IsMuted(),
		Width:     t.params.Width,
		Height:    t.params.Height,
		Simulcast: t.simulcasted,
	}
}

// this function assumes caller holds lock
func (t *MediaTrack) shouldStartWithBestQuality() bool {
	return len(t.subscribedTracks) < 10
}

// TODO: send for all downtracks from the source participant
// https://tools.ietf.org/html/rfc7941
func (t *MediaTrack) sendDownTrackBindingReports(sub types.Participant) {
	var sd []rtcp.SourceDescriptionChunk

	t.lock.RLock()
	subTrack := t.subscribedTracks[sub.ID()]
	t.lock.RUnlock()

	if subTrack == nil {
		return
	}

	chunks := subTrack.DownTrack().CreateSourceDescriptionChunks()
	if chunks == nil {
		return
	}
	sd = append(sd, chunks...)

	pkts := []rtcp.Packet{
		&rtcp.SourceDescription{Chunks: sd},
	}

	go func() {
		defer RecoverSilent()
		batch := pkts
		i := 0
		for {
			if err := sub.SubscriberPC().WriteRTCP(batch); err != nil {
				logger.Errorw("could not write RTCP", err)
				return
			}
			if i > 5 {
				return
			}
			i++
			time.Sleep(20 * time.Millisecond)
		}
	}()
}
