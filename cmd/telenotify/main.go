package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"time"

	// "github.com/inneroot/telenotify/internal/repository/memory"
	"github.com/inneroot/telenotify/internal/repository/pgrepo"
	"github.com/inneroot/telenotify/internal/telebot"
	"github.com/inneroot/telenotify/pkg/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := logger.SetLogger()

	repo, err := pgrepo.New(ctx, log, 5*time.Second)
	if err != nil {
		panic(fmt.Errorf("unable to init repo: %v", err.Error()))
	}
	// repo := memory.New()
	defer repo.Close()

	go func() {
		dbErr := telebot.Run(ctx, log, repo)
		if dbErr != nil {
			log.Error("failed to start telebot", "err", dbErr)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	log.Info("got interruption signal")
	// TODO graceful shutdown
	repo.Close()
	log.Info("telebot was shutdown gracefully")
}
