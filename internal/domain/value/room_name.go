package value

type RoomName struct {
	ValueObject[string]
}

func NewRoomName(v string) *RoomName {
	return &RoomName{NewValueObject(v)}
}
