package application

import (
	"context"

	"github.com/livekit/protocol/livekit"
)

type RoomManagementProxy interface {
	GetRoomList(ctx context.Context) ([]*livekit.Room, error)
}

type RoomManagementService struct {
	roomProxy RoomManagementProxy
}

func NewRoomManagementService(roomProxy RoomManagementProxy) *RoomManagementService {
	return &RoomManagementService{
		roomProxy: roomProxy,
	}
}

func (service *RoomManagementService) GetAllRooms(ctx context.Context) ([]*livekit.Room, error) {
	return service.roomProxy.GetRoomList(ctx)
}
