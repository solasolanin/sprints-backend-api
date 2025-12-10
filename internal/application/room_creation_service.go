package application

import (
	"context"
	"sprinta-backend-api/internal/domain/entity"
	"sprinta-backend-api/internal/domain/value"

	"github.com/livekit/protocol/livekit"
)

// CreateRoomメソッドを持つインターフェース（RoomManagementProxy用）
type IRoomCreationProxy interface {
    CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error)
}

// RegisterRoomメソッドを持つインターフェース（RoomRepository用）
type IRoomRepository interface {
    RegisterRoom(ctx context.Context, room entity.Room) error
}

type RoomCreationService struct {
    roomCreationProxy IRoomCreationProxy
    roomRepository    IRoomRepository
}

func NewRoomCreationService(roomCreationProxy IRoomCreationProxy, roomRepository IRoomRepository) *RoomCreationService {
    return &RoomCreationService{
        roomCreationProxy: roomCreationProxy,
        roomRepository:    roomRepository,
    }
}

func (s *RoomCreationService) CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error) {
	lkroom, err := s.roomCreationProxy.CreateRoom(ctx, roomName)
	if err != nil {
		return nil, err
	}

	room := entity.Room{
		Sid:              value.NewRoomSid(lkroom.Sid),
		Name:             value.NewRoomName(lkroom.Name),
		ParticipantCount: value.NewParticipantCount(lkroom.NumParticipants),
		Rank:             value.NewRoomRank("1"),
		CreatedAt:        value.NewRoomCreatedAtFromUnix(lkroom.CreationTime),
	}

	err = s.roomRepository.RegisterRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return lkroom, nil
}
