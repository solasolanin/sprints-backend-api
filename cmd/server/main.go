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

	proxy := proxy.NewRoomProxy(liveKitClient)
	service := application.NewRoomService(proxy)
	handler := handler.NewRoomHandler(service)

	r := gin.Default()

	// Define your routes here
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/room-list", handler.ListRooms)
	r.POST("/room", handler.CreateRoom)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
