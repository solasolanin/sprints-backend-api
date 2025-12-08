package request

type CreateRoomRequest struct {
	RoomName string `json:"roomName" binding:"required,max=10"`
}
