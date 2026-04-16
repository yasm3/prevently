package main

import (
	"github.com/yasm3/prevently/internal/api"
	"github.com/yasm3/prevently/internal/logger"
)

func main() {
	logger := logger.New()
	router := api.NewRouter(logger)
	server := api.NewServer(router, logger)
	server.Run()
}
