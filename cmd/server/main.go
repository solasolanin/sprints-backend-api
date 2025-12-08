package main

import (
	"log"
	"sprinta-backend-api/internal/api/handler"
	"sprinta-backend-api/internal/application"
	"sprinta-backend-api/internal/infrastructure/client"
	"sprinta-backend-api/internal/infrastructure/proxy"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func env_road() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// .envファイルの読み込み
	env_road()

	// LiveKitクライアントをアプリケーション起動時に一度だけ初期化
	liveKitClient, err := client.NewLiveKitClient()
	if err != nil {
		log.Fatalf("Failed to create LiveKit client: %v", err)
	}

	proxy := proxy.NewRoomManagementProxy(liveKitClient)
	service := application.NewRoomManagementService(proxy)
	handler := handler.NewRoomHandler(service)

	r := gin.Default()

	r.SetTrustedProxies([]string{
		// TODO 実際のCloud Load BalancerのIPレンジに変更する
		"35.191.0.0/16",
		"130.211.0.0/22",
	})

	r.GET("/room-list", handler.ListRooms)
	r.POST("/room", handler.CreateRoom)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
