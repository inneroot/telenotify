package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	httpRouter "github.com/inneroot/telenotify/internal/api/httprouter"
	notify_service "github.com/inneroot/telenotify/internal/service"
)

type HttpServer struct {
	log        *slog.Logger
	httpServer *http.Server
	port       int
}

func New(notifier notify_service.INotifier, port int, logger *slog.Logger) *HttpServer {
	ns := notify_service.New(notifier)
	router := httpRouter.New(ns)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	log := logger.With(slog.String("module", "httpserver"))
	return &HttpServer{
		log,
		httpServer,
		port,
	}

}

func (s *HttpServer) MustRunInGoRoutine() {
	go func() {
		if err := s.Run(); err != nil {
			os.Exit(1)
		}
	}()
}

func (s *HttpServer) Run() error {
	const op = "Run"
	log := s.log.With(slog.String("op", op))

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("http server listening", slog.Int("port", s.port))
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) {
	const op = "Stop"
	log := s.log.With(slog.String("op", op))
	log.Info("stopping grpc server")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}
}
