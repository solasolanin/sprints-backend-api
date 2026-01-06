package value

type NoticeType string

const (
	INFO  = NoticeType("INFO")
	NEWS  = NoticeType("NEWS")
	TIPS  = NoticeType("TIPS")
	EVENT = NoticeType("EVENT")
)

func NewNoticeType(s string) *NoticeType {
	t := NoticeType(s)
	return &t
}
