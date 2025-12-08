package request

type GetRoleTokenRequest struct {
	RoomName string `json:"roomName"`
	Identity string `json:"identity"`
}