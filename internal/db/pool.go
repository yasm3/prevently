package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() (*Queries, *pgxpool.Pool) {
	pool, err := pgxpool.New(
		context.Background(),
		"postgres://prevently:prevently@localhost:5432/prevently?sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	queries := New(pool)

	return queries, pool
}
