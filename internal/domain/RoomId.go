package domain

type RoomId struct {
	value string
}

func of(v string) RoomId {
	return RoomId{value: v}
}

func (r RoomId) Value() string {
	return r.value
}