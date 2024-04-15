package main

import (
	"context"
	"os"

	"time"

	// "github.com/inneroot/telenotify/internal/repository/memory"
	"github.com/inneroot/telenotify/internal/repository/pgrepo"
	"github.com/inneroot/telenotify/internal/telebot"
	"github.com/inneroot/telenotify/pkg/logger"
)

func main() {
	defaultCtx := context.Background()

	log := logger.SetLogger()

	repo, err := pgrepo.New(defaultCtx, log, 5*time.Second)
	if err != nil {
		panic("unable to init repo")
	}
	// repo := memory.New()
	defer repo.Close()

	dbErr := telebot.Run(defaultCtx, log, repo)
	if dbErr != nil {
		log.Error("failed to start telebot", "err", dbErr)
		os.Exit(1)
	}
}
