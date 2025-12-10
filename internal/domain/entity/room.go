package entity

import "sprinta-backend-api/internal/domain/value"

type Room struct {
	Sid  *value.RoomSid
	Name *value.RoomName
	ParticipantCount *value.ParticipantCount
	Rank *value.RoomRank
	CreatedAt *value.RoomCreatedAt
}
