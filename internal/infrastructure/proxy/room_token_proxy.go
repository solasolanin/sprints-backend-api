package proxy

import (
	"os"
	"time"

	"sprinta-backend-api/internal/config"

	"github.com/livekit/protocol/auth"
)

type RoomTokenProxy struct {
	token     string
	apiKey    string
	apiSecret string
}

func NewRoomTokenProxy(room string, identity string) (*RoomTokenProxy, error) {
	liveKitAPIKey := os.Getenv("LIVEKIT_API_KEY")
	liveKitAPISecret := os.Getenv("LIVEKIT_API_SECRET")

	at := auth.NewAccessToken(liveKitAPIKey, liveKitAPISecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.SetVideoGrant(grant).
		SetIdentity(identity).
		SetValidFor(time.Hour)

	token, err := at.ToJWT()
	if err != nil {
		return nil, &config.BussinessError{
			ErrorCode: 500,
			Message:   "failed to generate room JWT token",
		}
	}

	return &RoomTokenProxy{
		token:     token,
		apiKey:    liveKitAPIKey,
		apiSecret: liveKitAPISecret,
	}, nil
}

func (proxy *RoomTokenProxy) GetTokenClient() string {
	return proxy.token
}
