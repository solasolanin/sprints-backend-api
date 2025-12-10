package value

type RoomRank struct {
	ValueObject[string]
}

func NewRoomRank(v string) *RoomRank {
	return &RoomRank{NewValueObject(v)}
}
