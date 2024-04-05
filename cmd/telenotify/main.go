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

	postgres, dbErr := pg.NewPG(defaultCtx, log, connStr)
	if dbErr != nil {
		log.Error("failed to connect to db", dbErr)
		os.Exit(1)
	}
	defer postgres.Close()

	log.Debug(connStr)
	pgErr := postgres.WaitConnection(defaultCtx)
	if pgErr != nil {
		log.Error("failed to ping db", "err", pgErr)
		os.Exit(1)
	}
	log.Info("success: connected to db")

	tbErr := telebot.Run(defaultCtx, log)
	if tbErr != nil {
		log.Error("failed to start telebot", "err", dbErr)
		os.Exit(1)
	}
}
