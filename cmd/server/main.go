package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"sprinta-backend-api/internal/api/handler"
	"sprinta-backend-api/internal/application"
	"sprinta-backend-api/internal/config"
	"sprinta-backend-api/internal/infrastructure/client"
	"sprinta-backend-api/internal/infrastructure/proxy"
	"sprinta-backend-api/internal/infrastructure/repository"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func roadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := "5432"
	user := os.Getenv("DB_USER")
	password := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, &config.UnprocessableError{
			ErrorCode: 500,
			Message:   err.Error(),
		}
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func main() {
	// .envファイルの読み込み
	roadEnv()

	// LiveKitクライアントをアプリケーション起動時に一度だけ初期化
	liveKitClient, err := client.NewLiveKitClient()
	if err != nil {
		log.Fatalf("Failed to create LiveKit client: %v", err)
		return
	}

	// DB接続の初期化
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}

	// Firebaseアプリの初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Firebase Authクライアントの取得
	_, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	proxy := proxy.NewRoomManagementProxy(liveKitClient)
	roomManagementService := application.NewRoomManagementService(proxy)
	roomRepository := repository.NewRoomRepository(db)
	roomCreationService := application.NewRoomCreationService(proxy, roomRepository)
	roomTokenService := application.NewRoomTokenService()
	handler := handler.NewRoomHandler(roomManagementService, roomTokenService, roomCreationService)

	r := gin.Default()

	r.SetTrustedProxies([]string{
		// TODO 実際のCloud Load BalancerのIPレンジに変更する
		"35.191.0.0/16",
		"130.211.0.0/22",
	})

	r.GET("/room-list", handler.ListRooms)
	r.POST("/room", handler.CreateRoom)

	r.POST("/room-token", handler.GetRoomToken)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
