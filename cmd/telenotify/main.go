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
	grpcserver "github.com/inneroot/telenotify/internal/server"
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

	grpcServer := grpcserver.New(log, 5555)

	go func() {
		if err := grpcServer.Run(); err != nil {
			os.Exit(1)
		}
	}()
	go func() {
		dbErr := telebot.Run(ctx, log, repo)
		if dbErr != nil {
			log.Error("failed to start telebot", "err", dbErr)
			os.Exit(1)
		}
	}()

	<-ctx.Done() // graceful shutdown
	log.Info("got interruption signal")
	// repo.Close()
	grpcServer.Stop()
	log.Info("telebot was shutdown gracefully")
}
