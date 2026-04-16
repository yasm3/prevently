package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/domain"
	"github.com/yasm3/prevently/internal/http/dto"
	"github.com/yasm3/prevently/internal/http/middleware"
	"github.com/yasm3/prevently/internal/service"
)

type DeviceHandler struct {
	service *service.DeviceService
}

func NewDeviceHandler(s *service.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		service: s,
	}
}

func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	var body dto.CreateDeviceRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	u, exists := c.Get(middleware.UserContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, ResponseError{Error: "Unauthorized"})
		return
	}

	user, ok := u.(domain.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Invalid user in context"})
	}

	device, err := h.service.CreateDevice(
		c.Request.Context(),
		user.ID,
		body.Name,
		body.Type,
		body.Config,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to create device"})
		return
	}

	c.JSON(201, dto.CreateDeviceResponse{
		ID:     device.ID,
		UserID: device.UserID,
		Name:   device.Name,
		Type:   device.Type,
		Config: device.Config,
	})
}
