package client

import (
	"fmt"
	"os"

	lksdk "github.com/livekit/server-sdk-go/v2"
)

// LiveKitClient はLiveKitのクライアント操作をまとめた構造体です。
type LiveKitClient struct {
	roomClient *lksdk.RoomServiceClient
	apiKey     string
	apiSecret  string
}

// NewLiveKitClient はLiveKitClientを初期化します。
// サーバーURL、APIキー、APIシークレットは環境変数から取得します。
// 環境変数: LIVEKIT_SERVER_URL, LIVEKIT_API_KEY, LIVEKIT_API_SECRET
func NewLiveKitClient() (*LiveKitClient, error) {
	liveKitServerURL := os.Getenv("LIVEKIT_SERVER_URL")
	liveKitAPIKey := os.Getenv("LIVEKIT_API_KEY")
	liveKitAPISecret := os.Getenv("LIVEKIT_API_SECRET")

	if liveKitServerURL == "" {
		return nil, fmt.Errorf("environment variable LIVEKIT_SERVER_URL is not set")
	}
	if liveKitAPIKey == "" || liveKitAPISecret == "" {
		return nil, fmt.Errorf("environment variable LIVEKIT_API_KEY is not set")
	}

	roomClient := lksdk.NewRoomServiceClient(liveKitServerURL, liveKitAPIKey, liveKitAPISecret)

	return &LiveKitClient{
		roomClient: roomClient,
		apiKey:     liveKitAPIKey,
		apiSecret:  liveKitAPISecret,
	}, nil
}

// RoomServiceClient は初期化されたRoomServiceClientを返します。
func (lk *LiveKitClient) RoomServiceClient() *lksdk.RoomServiceClient {
	return lk.roomClient
}
