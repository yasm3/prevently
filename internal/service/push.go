package service

import (
	"context"

	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/domain"
)

type PushService struct {
	db *db.Queries
}

func NewPushService(q *db.Queries) *PushService {
	return &PushService{
		db: q,
	}
}

func (s *PushService) CreatePush(ctx context.Context, userId string, message string) (domain.Push, error) {
	row, err := s.db.CreatePush(ctx, db.CreatePushParams{
		UserID:  userId,
		Message: message,
	})
	if err != nil {
		return domain.Push{}, err
	}

	return domain.Push{
		ID:        row.ID,
		UserID:    row.UserID,
		Message:   row.Message,
		Status:    row.Status,
		Attempts:  int(row.Attempts),
		LastError: row.LastError.String,
		CreatedAt: row.CreatedAt.Time,
		SentAt:    row.SentAt.Time,
	}, nil
}
