package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yasm3/prevently/internal/db"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, uuid pgtype.UUID) (db.User, error) {
	return s.db.GetUserByID(ctx, uuid)
}
