package dto

import (
	"encoding/json"

	"github.com/yasm3/prevently/internal/domain"
)

type CreateDeviceRequest struct {
	Name   string            `json:"name" binding:"required"`
	Type   domain.DeviceType `json:"type" binding:"required,oneof=discord"`
	Config json.RawMessage   `json:"config" binding:"required"`
}

type CreateDeviceResponse struct {
	ID     string            `json:"id"`
	UserID string            `json:"user_id"`
	Name   string            `json:"name"`
	Type   domain.DeviceType `json:"type"`
	Config json.RawMessage   `json:"config"`
}
