package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/domain"
	"github.com/yasm3/prevently/internal/http/dto"
	"github.com/yasm3/prevently/internal/http/middleware"
	"github.com/yasm3/prevently/internal/service"
)

type PushHandler struct {
	service *service.PushService
}

func NewPushHandler(s *service.PushService) *PushHandler {
	return &PushHandler{
		service: s,
	}
}

func (h *PushHandler) SendPush(c *gin.Context) {
	var body dto.CreatePushRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	if body.Message == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Error: "Message must not be empty"})
	}

	u, exists := c.Get(middleware.UserContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, ResponseError{Error: "Unauthorized"})
		return
	}

	user, ok := u.(domain.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Invalid user in context"})
		return
	}

	_, err := h.service.CreatePush(c.Request.Context(), user.ID, body.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to push"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Your message will be delivered within seconds",
	})
}
