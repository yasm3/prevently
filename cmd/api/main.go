package main

import (
	"github.com/yasm3/prevently/internal/api"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
)

func main() {
	logger := logger.New()

	queries, pool := db.NewDB()
	defer pool.Close()

	router := api.NewRouter(logger)

	server := api.NewServer(router, queries, logger)
	server.Run()
}
