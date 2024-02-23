package main

import (
	"context"
	"os"

	telebot "github.com/inneroot/telenotify/internal/telebot"
	"github.com/inneroot/telenotify/pkg/logger"
)

func main() {
	defaultCtx := context.Background()

	log := logger.SetLogger()

	err := telebot.Run(defaultCtx)
	if err != nil {
		log.Error("failed to start telebot", logger.Err(err))
		os.Exit(1)
	}
}
