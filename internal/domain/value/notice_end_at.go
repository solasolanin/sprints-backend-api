package value

import (
	"encoding/json"
	"time"
)

type NoticeEndAt struct {
	ValueObject[time.Time]
}

func NewNoticeEndAt(v time.Time) *NoticeEndAt {
	return &NoticeEndAt{NewValueObject(v)}
}

func NewNoticeEndAtNow() *NoticeEndAt {
	return &NoticeEndAt{NewValueObject(time.Now())}
}

func NewNoticeEndAtFromUnix(v int64) *NoticeEndAt {
	return &NoticeEndAt{NewValueObject(time.Unix(v, 0))}
}

func (r *NoticeEndAt) Unix() int64 {
	return r.Value().Unix()
}

func (r *NoticeEndAt) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	return json.Marshal(r.Value().Format(time.DateTime))
}
