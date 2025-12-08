package handler

import (
	http "net/http"
	"sprinta-backend-api/internal/api/request"
	"sprinta-backend-api/internal/application"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomService *application.RoomService
}

func NewRoomHandler(roomService *application.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) ListRooms(ctx *gin.Context) {
	rooms, err := h.roomService.GetAllRooms(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	var req request.CreateRoomRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.roomService.CreateRoom(ctx, req.RoomName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, room)
}
