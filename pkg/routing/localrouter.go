package routing

import (
	"sync"
	"time"

	"github.com/livekit/livekit-server/proto/livekit"
)

// a router of messages on the same node, basic implementation for local testing
type LocalRouter struct {
	currentNode LocalNode
	lock        sync.Mutex
	// channels for each participant
	requestChannels  map[string]*MessageChannel
	responseChannels map[string]*MessageChannel
	onNewParticipant ParticipantCallback
}

func NewLocalRouter(currentNode LocalNode) *LocalRouter {
	return &LocalRouter{
		currentNode:      currentNode,
		requestChannels:  make(map[string]*MessageChannel),
		responseChannels: make(map[string]*MessageChannel),
	}
}

func (r *LocalRouter) GetNodeForRoom(roomName string) (string, error) {
	return r.currentNode.Id, nil
}

func (r *LocalRouter) SetNodeForRoom(roomName string, nodeId string) error {
	return nil
}

func (r *LocalRouter) ClearRoomState(roomName string) error {
	// do nothing
	return nil
}

func (r *LocalRouter) RegisterNode() error {
	return nil
}

func (r *LocalRouter) UnregisterNode() error {
	return nil
}

func (r *LocalRouter) GetNode(nodeId string) (*livekit.Node, error) {
	if nodeId == r.currentNode.Id {
		return r.currentNode, nil
	}
	return nil, ErrNotFound
}

func (r *LocalRouter) ListNodes() ([]*livekit.Node, error) {
	return []*livekit.Node{
		r.currentNode,
	}, nil
}

func (r *LocalRouter) StartParticipantSignal(roomName, participantId, participantName string) error {
	// treat it as a new participant connecting
	if r.onNewParticipant == nil {
		return ErrHandlerNotDefined
	}
	r.onNewParticipant(
		roomName,
		participantId,
		participantName,
		// request source
		r.getOrCreateMessageChannel(r.requestChannels, participantId),
		// response sink
		r.getOrCreateMessageChannel(r.responseChannels, participantId),
	)
	return nil
}

// for a local router, sink and source are pointing to the same spot
func (r *LocalRouter) GetRequestSink(participantId string) (MessageSink, error) {
	return r.getOrCreateMessageChannel(r.requestChannels, participantId), nil
}

func (r *LocalRouter) GetResponseSource(participantId string) (MessageSource, error) {
	return r.getOrCreateMessageChannel(r.responseChannels, participantId), nil
}

func (r *LocalRouter) OnNewParticipantRTC(callback ParticipantCallback) {
	r.onNewParticipant = callback
}

func (r *LocalRouter) Start() error {
	go r.statsWorker()
	// on local routers, Start doesn't do anything, websocket connections initiate the connections
	return nil
}

func (r *LocalRouter) Stop() {
}

func (r *LocalRouter) statsWorker() {
	for {
		// update every 10 seconds
		<-time.After(statsUpdateInterval)
		r.currentNode.Stats.UpdatedAt = time.Now().Unix()
	}
}

func (r *LocalRouter) getOrCreateMessageChannel(target map[string]*MessageChannel, participantId string) *MessageChannel {
	r.lock.Lock()
	defer r.lock.Unlock()
	mc := target[participantId]

	if mc != nil {
		return mc
	}

	mc = NewMessageChannel()
	mc.OnClose(func() {
		r.lock.Lock()
		delete(target, participantId)
		r.lock.Unlock()
	})
	target[participantId] = mc

	return mc
}
