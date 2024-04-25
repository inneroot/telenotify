package main

import (
	"context"
	"fmt"
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
	grpcServer.MustRunInGoRoutine()
	defer grpcServer.Stop()

	bot := telebot.MustInit(ctx, log, repo)
	bot.Run()
	defer bot.Stop()

	<-ctx.Done() // graceful shutdown with deferred functions
	log.Info("telebot will be shutdown gracefully")
}
