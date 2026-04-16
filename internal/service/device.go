package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/domain"
)

type DeviceService struct {
	db *db.Queries
}

func NewDeviceService(q *db.Queries) *DeviceService {
	return &DeviceService{
		db: q,
	}
}

func (s *DeviceService) validateConfig(t domain.DeviceType, raw json.RawMessage) error {
	switch t {
	case domain.Discord:
		var cfg domain.DiscordConfig
		if err := json.Unmarshal(raw, &cfg); err != nil {
			return errors.New("Invalid Discord coonfig json")
		}
		if cfg.WebhookURL == "" {
			return errors.New("webhook_url is required")
		}
	default:
		return errors.New("Unsupported device type")
	}

	return nil
}

func (s *DeviceService) CreateDevice(
	ctx context.Context,
	userID string,
	name string,
	deviceType domain.DeviceType,
	config []byte,
) (domain.Device, error) {

	if err := s.validateConfig(deviceType, config); err != nil {
		return domain.Device{}, err
	}

	row, err := s.db.CreateDevice(ctx, db.CreateDeviceParams{
		UserID: userID,
		Name:   name,
		Type:   string(deviceType),
		Config: config,
	})
	if err != nil {
		return domain.Device{}, err
	}

	return domain.Device{
		ID:     row.ID,
		UserID: row.UserID,
		Name:   row.Name,
		Type:   domain.DeviceType(row.Type),
		Config: row.Config,
	}, nil
}
