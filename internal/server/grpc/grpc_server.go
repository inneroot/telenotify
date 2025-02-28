package grpcserver

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/inneroot/telenotify/internal/api/grpchandler"
	notify_service "github.com/inneroot/telenotify/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(notifier notify_service.INotifier, port int, logger *slog.Logger) *GRPCServer {
	gRPCServer := grpc.NewServer()
	ns := notify_service.New(notifier)
	serverApi := grpchandler.New(ns)
	grpchandler.RegisterNotifyServiceServer(gRPCServer, serverApi)
	reflection.Register(gRPCServer)
	log := logger.With(slog.String("module", "grpcserver"))
	return &GRPCServer{
		log,
		gRPCServer,
		port,
	}
}

func (s *GRPCServer) MustRunInGoRoutine() {
	go func() {
		if err := s.Run(); err != nil {
			s.log.Error(err.Error())
			os.Exit(1)
		}
	}()
}

func (s *GRPCServer) Run() error {
	const op = "Run"
	log := s.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpc server listening", slog.String("addr", l.Addr().String()), slog.Int("port", s.port))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) {
	const op = "Stop"
	s.log.With(slog.String("op", op)).Info("stopping grpc server")
	s.gRPCServer.GracefulStop()
}
