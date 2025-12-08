package application

import (
	"context"

	"github.com/livekit/protocol/livekit"
)

type RoomProxy interface {
	CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error)
	GetRoomList(ctx context.Context) ([]*livekit.Room, error)
}

type RoomService struct {
	roomProxy RoomProxy
}

func NewRoomService(roomProxy RoomProxy) *RoomService {
	return &RoomService{
		roomProxy: roomProxy,
	}
}

func (service *RoomService) GetAllRooms(ctx context.Context) ([]*livekit.Room, error) {
	return service.roomProxy.GetRoomList(ctx)
}

func (service *RoomService) CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error) {
	return service.roomProxy.CreateRoom(ctx, roomName)
}
