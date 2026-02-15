package handler

import (
	http "net/http"
	"sprinta-backend-api/internal/api/request"
	"sprinta-backend-api/internal/application"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomManagementService *application.RoomManagementService
	roomTokenService      *application.RoomTokenService
	roomCreationService   *application.RoomCreationService
	roomWebhookService    *application.RoomWebhookService
}

func NewRoomHandler(
	roomManagementService *application.RoomManagementService,
	roomTokenService *application.RoomTokenService,
	roomCreationService *application.RoomCreationService,
	roomWebhookService *application.RoomWebhookService,
) *RoomHandler {
	return &RoomHandler{
		roomManagementService: roomManagementService,
		roomTokenService:      roomTokenService,
		roomCreationService:   roomCreationService,
		roomWebhookService:    roomWebhookService,
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

	room, err := h.roomCreationService.CreateRoom(ctx, req.RoomName)
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

func (h *RoomHandler) LiveKitWebhook(ctx *gin.Context) {
	var payload map[string]interface{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// try to detect event type
	var eventType string
	if v, ok := payload["type"].(string); ok {
		eventType = v
	}
	if ev, ok := payload["event"].(map[string]interface{}); ok {
		if t, ok2 := ev["type"].(string); ok2 {
			eventType = t
		}
	}

	// only handle participant left/disconnected events
	if eventType == "" || !(strings.Contains(strings.ToLower(eventType), "participant") && (strings.Contains(strings.ToLower(eventType), "left") || strings.Contains(strings.ToLower(eventType), "disconnect") || strings.Contains(strings.ToLower(eventType), "disconnected"))) {
		ctx.Status(http.StatusOK)
		return
	}

	// extract room sid
	var sid string
	if roomObj, ok := payload["room"].(map[string]interface{}); ok {
		if s, ok2 := roomObj["sid"].(string); ok2 {
			sid = s
		}
	}
	if sid == "" {
		if s, ok := payload["room_sid"].(string); ok {
			sid = s
		}
	}

	if sid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room sid not found in payload"})
		return
	}

	if err := h.roomWebhookService.HandleParticipantLeft(ctx.Request.Context(), sid); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
