package main

import (
	"os"

	"github.com/yasm3/prevently/internal/api"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
)

func main() {
	logger := logger.New()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://prevently:prevently@localhost:5432/prevently?sslmode=disable"
	}

	queries, pool := db.NewDB(dbUrl)
	defer pool.Close()

	server := api.NewServer(queries, logger)

	server.Run()
}
