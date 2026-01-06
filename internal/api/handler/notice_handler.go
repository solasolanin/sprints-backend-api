package handler

import (
	http "net/http"
	"sprinta-backend-api/internal/application"

	"github.com/gin-gonic/gin"
)

type NoticeHandler struct{}

func NewNoticeHandler() *NoticeHandler {
	return &NoticeHandler{}
}

// ListNotices calls application.GetNoticeService and returns notices as JSON.
// No request parameters are required.
func (h *NoticeHandler) ListNotices(ctx *gin.Context) {
	notices, err := application.GetNoticeService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notices)
}
