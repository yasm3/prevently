package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(dbUrl string) (*Queries, *pgxpool.Pool) {
	pool, err := pgxpool.New(
		context.Background(),
		dbUrl,
	)
	if err != nil {
		panic(err)
	}

	queries := New(pool)

	return queries, pool
}
