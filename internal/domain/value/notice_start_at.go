package value

import (
	"encoding/json"
	"time"
)

type NoticeStartAt struct {
	ValueObject[time.Time]
}

func NewNoticeStartAt(v time.Time) *NoticeStartAt {
	return &NoticeStartAt{NewValueObject(v)}
}

func NewNoticeStartAtNow() *NoticeStartAt {
	return &NoticeStartAt{NewValueObject(time.Now())}
}

func NewNoticeStartAtFromUnix(v int64) *NoticeStartAt {
	return &NoticeStartAt{NewValueObject(time.Unix(v, 0))}
}

func (r *NoticeStartAt) Unix() int64 {
	return r.Value().Unix()
}

func (r *NoticeStartAt) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	return json.Marshal(r.Value().Format(time.DateTime))
}
