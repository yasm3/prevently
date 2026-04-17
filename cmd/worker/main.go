package main

import (
	"context"
	"os"
	"time"

	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
	"github.com/yasm3/prevently/internal/service"
)

func main() {
	ctx := context.Background()
	logger := logger.New()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://prevently:prevently@localhost:5432/prevently?sslmode=disable"
	}

	queries, pool := db.NewDB(dbUrl)
	defer pool.Close()

	pushService := service.NewPushService(queries)

	logger.Info("Worker started...")

	for {
		pushes, err := pushService.ClaimPendingPushes(ctx, 10)
		if err != nil {
			logger.Info(err.Error())
			continue
		}

		for _, p := range pushes {
			if err := pushService.ProcessPush(ctx, p); err != nil {
				logger.Info(err.Error())
			}
		}

		time.Sleep(time.Second * 1)
	}

}
