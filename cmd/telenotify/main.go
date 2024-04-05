package main

import (
	"context"
	"os"

	"github.com/inneroot/telenotify/internal/config"
	"github.com/inneroot/telenotify/internal/telebot"
	"github.com/inneroot/telenotify/pkg/logger"
	"github.com/inneroot/telenotify/pkg/pg"
)

func main() {
	defaultCtx := context.Background()

	log := logger.SetLogger()

	connStr := config.GetPGConnectionString()

	postgres, dbErr := pg.NewPG(defaultCtx, connStr)
	if dbErr != nil {
		log.Error("failed to connect to db", dbErr)
		os.Exit(1)
	}

	pingErr := postgres.Ping(defaultCtx)
	if pingErr != nil {
		log.Error("failed to ping db", dbErr)
		os.Exit(1)
	}

	tbErr := telebot.Run(defaultCtx)
	if tbErr != nil {
		log.Error("failed to start telebot", dbErr)
		os.Exit(1)
	}
}
