package application

import (
	"sprinta-backend-api/internal/domain/entity"
	"sprinta-backend-api/internal/domain/value"
	"time"
)

func GetNoticeService() ([]*entity.PublicNotice, error) {
	startAt, err := time.Parse(time.DateTime, "2025-12-10 00:00:00")
	if err != nil {
		return nil, err
	}
	endAt, err := time.Parse(time.DateTime, "2026-12-10 00:00:00")
	if err != nil {
		return nil, err
	}

	notices := []*entity.PublicNotice{
		{
			Id:      value.NewNoticeId("0001"),
			Title:   value.NewNoticeTitle("サービス開始"),
			Type:    value.NewNoticeType("info"),
			StartAt: value.NewNoticeStartAt(startAt),
			EndAt:   value.NewNoticeEndAt(endAt),
		},
		{
			Id:      value.NewNoticeId("0002"),
			Title:   value.NewNoticeTitle("新BGM\n追加機能"),
			Type:    value.NewNoticeType("new_feature"),
			StartAt: value.NewNoticeStartAt(startAt),
			EndAt:   value.NewNoticeEndAt(endAt),
		},
		{
			Id:		value.NewNoticeId("0003"),
			Title: 	value.NewNoticeTitle("週末Sprint"),
			Type:	value.NewNoticeType("event"),
			StartAt: value.NewNoticeStartAt(startAt),
			EndAt:   value.NewNoticeEndAt(endAt),
		},
	}

	print(&notices)

	return notices, nil
}
