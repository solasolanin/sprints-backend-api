package entity

import "sprinta-backend-api/internal/domain/value"

type PublicNotice struct {
	Id      *value.NoticeId      `json:"id"`
	Title   *value.NoticeTitle   `json:"title"`
	Type    *value.NoticeType    `json:"type"`
	StartAt *value.NoticeStartAt `json:"startAt"`
	EndAt   *value.NoticeEndAt   `json:"endAt"`
}
