package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/livekit-server/pkg/utils"
	"github.com/livekit/livekit-server/proto/livekit"
)

const (
	// hash of room_name => Room proto
	RoomsKey = "rooms"

	// hash of room_id => room name
	RoomIdMap = "room_id_map"

	// hash of participant_name => participant_id
	// a key for each room, with expiration
	RoomParticipantMapPrefix = "participant_map:room:"

	participantMappingTTL = 24 * time.Hour
)

type RedisRoomStore struct {
	rc  *redis.Client
	ctx context.Context
}

func NewRedisRoomStore(rc *redis.Client) *RedisRoomStore {
	return &RedisRoomStore{
		ctx: context.Background(),
		rc:  rc,
	}
}

func (p *RedisRoomStore) CreateRoom(room *livekit.Room) error {
	if room.CreationTime == 0 {
		room.CreationTime = time.Now().Unix()
	}

	data, err := proto.Marshal(room)
	if err != nil {
		return err
	}

	pp := p.rc.Pipeline()
	pp.HSet(p.ctx, RoomIdMap, room.Sid, room.Name)
	pp.HSet(p.ctx, RoomsKey, room.Name, data)

	if _, err = pp.Exec(p.ctx); err != nil {
		return errors.Wrap(err, "could not create room")
	}
	return nil
}

func (p *RedisRoomStore) GetRoom(idOrName string) (*livekit.Room, error) {
	// see if matches any ids
	name, err := p.rc.HGet(p.ctx, RoomIdMap, idOrName).Result()
	if err != nil {
		name = idOrName
	}

	data, err := p.rc.HGet(p.ctx, RoomsKey, name).Result()
	if err != nil {
		if err == redis.Nil {
			err = ErrRoomNotFound
		}
		return nil, err
	}

	room := livekit.Room{}
	err = proto.Unmarshal([]byte(data), &room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (p *RedisRoomStore) ListRooms() ([]*livekit.Room, error) {
	items, err := p.rc.HVals(p.ctx, RoomsKey).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "could not get rooms")
	}

	rooms := make([]*livekit.Room, 0, len(items))

	for _, item := range items {
		room := livekit.Room{}
		err := proto.Unmarshal([]byte(item), &room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (p *RedisRoomStore) DeleteRoom(idOrName string) error {
	room, err := p.GetRoom(idOrName)
	if err == ErrRoomNotFound {
		return nil
	} else if err != nil {
		return err
	}

	pp := p.rc.Pipeline()
	pp.HDel(p.ctx, RoomIdMap, room.Sid)
	pp.HDel(p.ctx, RoomsKey, room.Name)
	pp.HDel(p.ctx, RoomParticipantMapPrefix+room.Name)

	_, err = pp.Exec(p.ctx)
	return err
}

func (p *RedisRoomStore) GetParticipantId(room, name string) (string, error) {
	key := RoomParticipantMapPrefix + room

	pId, err := p.rc.HGet(p.ctx, key, name).Result()
	if err == redis.Nil {
		// create
		pId = utils.NewGuid(utils.ParticipantPrefix)
		pp := p.rc.Pipeline()
		pp.HSet(p.ctx, key, name, pId)
		pp.Expire(p.ctx, key, participantMappingTTL)
		_, err = pp.Exec(p.ctx)
	} else if err != nil {
		return "", err
	}

	return pId, err
}
