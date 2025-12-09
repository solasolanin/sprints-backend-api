package request

type GetRoomTokenRequest struct {
	RoomName string `json:"roomName" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}