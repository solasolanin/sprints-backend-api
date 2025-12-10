package value

type RoomSid struct {
	ValueObject[string]
}

func NewRoomSid(v string) *RoomSid {
	return &RoomSid{NewValueObject(v)}
}