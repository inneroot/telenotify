package main

import (
	"context"
	"os/signal"
	"syscall"

	"time"

	// "github.com/inneroot/telenotify/internal/repository/memory"
	"github.com/inneroot/telenotify/internal/config"
	"github.com/inneroot/telenotify/internal/repository/pgrepo"
	grpcserver "github.com/inneroot/telenotify/internal/server/grpc"
	httpserver "github.com/inneroot/telenotify/internal/server/http"
	"github.com/inneroot/telenotify/internal/telebot"
	"github.com/inneroot/telenotify/pkg/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := logger.SetLogger()

	repo := pgrepo.MustInit(ctx, log, 5*time.Second)
	// repo := memory.New()
	defer repo.Close()
	serverPorts := config.GetServerPorts()

	log.Info("starting telegram bot")
	bot := telebot.MustInit(ctx, log, repo)
	bot.Run()
	defer bot.Stop()

	grpcServer := grpcserver.New(bot, serverPorts.GrpcPort, log)
	grpcServer.MustRunInGoRoutine()
	defer grpcServer.Stop(ctx)

	httpServer := httpserver.New(bot, serverPorts.HttpPort, log)
	httpServer.MustRunInGoRoutine()
	defer httpServer.Stop(ctx)

	<-ctx.Done() // graceful shutdown with deferred functions
	log.Info("telebot will be shutdown gracefully")
}
