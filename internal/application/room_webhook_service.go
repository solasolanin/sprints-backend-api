package application

import (
    "context"
)

// IRoomRepositoryUpdater provides repository methods needed by webhook service
type IRoomRepositoryUpdater interface {
    DecrementParticipantCountBySid(ctx context.Context, sid string) error
}

type RoomWebhookService struct {
    repo IRoomRepositoryUpdater
}

func NewRoomWebhookService(repo IRoomRepositoryUpdater) *RoomWebhookService {
    return &RoomWebhookService{repo: repo}
}

func (s *RoomWebhookService) HandleParticipantLeft(ctx context.Context, sid string) error {
    return s.repo.DecrementParticipantCountBySid(ctx, sid)
}
