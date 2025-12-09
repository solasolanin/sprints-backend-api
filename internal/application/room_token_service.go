package application

import "sprinta-backend-api/internal/infrastructure/proxy"

type RoomTokenService struct{}

func NewRoomTokenService() *RoomTokenService {
	return &RoomTokenService{}
}

func (service *RoomTokenService) GetRoomToken(room string, identity string) (string, error) {
	proxyInstance, err := proxy.NewRoomTokenProxy(room, identity)
	if err != nil {
		return "", err
	}
	return proxyInstance.GetTokenClient(), nil
}
