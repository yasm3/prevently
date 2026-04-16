package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/domain"
	"github.com/yasm3/prevently/internal/security"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

type CreateUserResponse struct {
	Email  string
	APIKey string
}

func (s *UserService) GetUserByID(ctx context.Context, uuid pgtype.UUID) (db.User, error) {
	return s.db.GetUserByID(ctx, uuid)
}

func (s *UserService) CreateUser(ctx context.Context, email string) (domain.User, string, error) {
	apiKey, err := security.GenerateUserAPIKey()
	if err != nil {
		return domain.User{}, "", err
	}

	user, err := s.db.CreateUser(ctx, db.CreateUserParams{
		Email:  email,
		ApiKey: security.HashAPIKey(apiKey),
	})
	if err != nil {
		return domain.User{}, "", err
	}

	return domain.User{
		ID:    user.ID.String(),
		Email: user.Email,
	}, apiKey, nil
}
