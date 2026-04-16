package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (h *UserHandler) GetUser(c *gin.Context) {
	// fake uuid test
	u, _ := uuid.NewV4()
	var pgid pgtype.UUID
	_ = pgid.Scan(u.String())

	user, err := h.service.GetUserByID(c.Request.Context(), pgid)
	if err != nil {
		c.JSON(404, ResponseError{Error: err.Error()})
		return
	}

	c.JSON(200, user)
}
