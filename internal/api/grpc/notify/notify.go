package notify

import (
	"context"
	"log/slog"

	telebot_notify "github.com/inneroot/telenotify/pkg/api/grpc"
	"google.golang.org/grpc"
)

type grpcServer struct {
	telebot_notify.UnimplementedNotifyServiceServer
}

func (s *grpcServer) Notify(ctx context.Context, req *telebot_notify.NotifyRequest) (*telebot_notify.NotifyResponse, error) {
	slog.Info("notify", slog.String("message", req.Message))
	return &telebot_notify.NotifyResponse{}, nil
}

func Notify(gRPC *grpc.Server) {
	telebot_notify.RegisterNotifyServiceServer(gRPC, &grpcServer{})
}
