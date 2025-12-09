package handler

import (
	http "net/http"
	"sprinta-backend-api/internal/api/request"
	"sprinta-backend-api/internal/application"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomManagementService *application.RoomManagementService
	roomTokenService      *application.RoomTokenService
}

func NewRoomHandler(
	roomManagementService *application.RoomManagementService,
	roomTokenService *application.RoomTokenService) *RoomHandler {
	return &RoomHandler{
		roomManagementService: roomManagementService,
		roomTokenService:      roomTokenService,
	}
}

func (h *RoomHandler) ListRooms(ctx *gin.Context) {
	rooms, err := h.roomManagementService.GetAllRooms(ctx)
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

	room, err := h.roomManagementService.CreateRoom(ctx, req.RoomName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, room)
}

func (h *RoomHandler) GetRoomToken(ctx *gin.Context) {
	var req request.GetRoomTokenRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.roomTokenService.GetRoomToken(req.RoomName, req.Identity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}	