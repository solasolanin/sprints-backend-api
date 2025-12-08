package proxy

import (
	"context"
	"sprinta-backend-api/internal/infrastructure/client"

	"github.com/livekit/protocol/livekit"
)

type RoomManagementProxy struct {
	client *client.LiveKitClient
}

func NewRoomManagementProxy(lkClient *client.LiveKitClient) *RoomManagementProxy {
	return &RoomManagementProxy{
		client: lkClient,
	}
}

func (roomProxy *RoomManagementProxy) CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error) {
	req := &livekit.CreateRoomRequest{
		Name: roomName,
	}
	res, err := roomProxy.client.RoomServiceClient().CreateRoom(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (roomProxy *RoomManagementProxy) GetRoomList(ctx context.Context) ([]*livekit.Room, error) {
	req := &livekit.ListRoomsRequest{}

	res, err := roomProxy.client.RoomServiceClient().ListRooms(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Rooms, nil
}
