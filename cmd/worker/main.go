package main

import (
	"context"
	"time"

	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
	"github.com/yasm3/prevently/internal/service"
)

func main() {
	ctx := context.Background()
	logger := logger.New()

	db, pool := db.NewDB()
	defer pool.Close()

	pushService := service.NewPushService(db)

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
