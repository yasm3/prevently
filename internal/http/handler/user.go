package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/domain"
	"github.com/yasm3/prevently/internal/http/dto"
	"github.com/yasm3/prevently/internal/http/middleware"
	"github.com/yasm3/prevently/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var body dto.CreateUserSchema
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(500, ResponseError{Error: err.Error()})
		return
	}

	body.Email = strings.TrimSpace(strings.ToLower(body.Email))

	user, apiKey, err := h.service.CreateUser(c.Request.Context(), body.Email)
	if err != nil {
		c.JSON(500, ResponseError{Error: err.Error()})
		return
	}

	c.JSON(201, dto.CreateUserResponse{
		ID:     user.ID,
		Email:  user.Email,
		APIKey: apiKey,
	})
}

func (h *UserHandler) GetMe(c *gin.Context) {
	u, exists := (c.Get(middleware.UserContextKey))
	if !exists {
		c.JSON(500, ResponseError{Error: "User not in context"})
		return
	}

	user, ok := u.(domain.User)
	if !ok {
		c.JSON(500, ResponseError{Error: "Invalid user type"})
		return
	}

	c.JSON(200, dto.GetUserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
