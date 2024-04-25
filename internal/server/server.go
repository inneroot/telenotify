package grpcserver

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	grpcNotify "github.com/inneroot/telenotify/internal/api/grpc/notify"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(logger *slog.Logger,
	port int) *GRPCServer {
	gRPCServer := grpc.NewServer()
	grpcNotify.Notify(gRPCServer)
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
	log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpcServer started")
	return nil
}

func (s *GRPCServer) Stop() {
	const op = "Stop"
	s.log.With(slog.String("op", op)).Info("stopping grpc server")
	s.gRPCServer.GracefulStop()
}
