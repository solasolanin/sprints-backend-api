package value

import "encoding/json"

type NoticeId struct {
	ValueObject[string]
}

func NewNoticeId(v string) *NoticeId {
	return &NoticeId{NewValueObject(v)}
}

func (n *NoticeId) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value())
}
