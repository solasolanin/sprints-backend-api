package value

import "encoding/json"

// お知らせタイトルを表すValue Object
type NoticeTitle struct {
	ValueObject[string]
}

func NewNoticeTitle(v string) *NoticeTitle {
	return &NoticeTitle{NewValueObject(v)}
}

func (n *NoticeTitle) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value())
}
