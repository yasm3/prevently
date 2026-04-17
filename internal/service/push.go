package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
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

func (s *PushService) ClaimPendingPushes(ctx context.Context, limit int) ([]domain.Push, error) {
	rows, err := s.db.ClaimPendingPushes(ctx, int32(limit))
	if err != nil {
		return []domain.Push{}, err
	}

	pushes := make([]domain.Push, 0, len(rows))
	for _, p := range rows {
		pushes = append(pushes, domain.Push{
			ID:        p.ID,
			UserID:    p.UserID,
			Message:   p.Message,
			Status:    p.Status,
			Attempts:  int(p.Attempts),
			LastError: p.LastError.String,
			CreatedAt: p.CreatedAt.Time,
			SentAt:    p.SentAt.Time,
		})
	}

	return pushes, nil
}

func postJSON(url string, body any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New("bad status")
	}

	return nil
}

func sendToDevice(d db.Device, message string) error {
	switch d.Type {
	case "discord":
		var cfg domain.DiscordConfig
		if err := json.Unmarshal(d.Config, &cfg); err != nil {
			return err
		}
		payload := map[string]string{
			"content": message,
		}
		return postJSON(cfg.WebhookURL, payload)
	default:
		return errors.New("Unsupported device type: " + d.Type)
	}
}

func (s *PushService) ProcessPush(ctx context.Context, p domain.Push) error {
	devices, err := s.db.ListDevicesByUser(ctx, p.UserID)
	if err != nil {
		_, _ = s.db.MarkPushFailed(ctx, db.MarkPushFailedParams{
			ID:        p.ID,
			LastError: pgtype.Text{String: err.Error()},
		})
		return err
	}

	var hasError bool
	for _, d := range devices {
		err := sendToDevice(d, p.Message)
		if err != nil {
			hasError = true
			log.Printf("Failed sending to device %s: %v", d.ID, err)
		}
	}

	if hasError {
		_, err := s.db.MarkPushFailed(ctx, db.MarkPushFailedParams{
			ID:        p.ID,
			LastError: pgtype.Text{String: "One or more devices failed"},
		})
		return err
	}

	s.db.MarkPushSent(ctx, p.ID)
	return nil
}
