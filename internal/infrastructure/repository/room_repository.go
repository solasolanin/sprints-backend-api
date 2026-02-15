package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sprinta-backend-api/internal/domain/entity"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) RegisterRoom(ctx context.Context, room entity.Room) error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := "INSERT INTO rooms (sid, name, participant_count, rank, created_at, delete_flg) VALUES ($1, $2, $3, $4, $5, false)"
	res, err := r.db.ExecContext(
		ctx,
		query,
		room.Sid.Value(),
		room.Name.Value(),
		room.ParticipantCount.Value(),
		room.Rank.Value(),
		room.CreatedAt.Value(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert room: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}
