package value

import "time"

type RoomCreatedAt struct {
	ValueObject[time.Time]
}

func NewRoomCreatedAt(v time.Time) *RoomCreatedAt {
	return &RoomCreatedAt{NewValueObject(v)}
}

func NewRoomCreatedAtNow() *RoomCreatedAt {
	return &RoomCreatedAt{NewValueObject(time.Now())}
}

func NewRoomCreatedAtFromUnix(v int64) *RoomCreatedAt {
	return &RoomCreatedAt{NewValueObject(time.Unix(v, 0))}
}

func (r *RoomCreatedAt) Unix() int64 {
	return r.Value().Unix()
}