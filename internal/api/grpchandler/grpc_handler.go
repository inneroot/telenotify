package grpchandler

import (
	"context"
	"log/slog"

	notify_service "github.com/inneroot/telenotify/internal/service"
	telebot_notify "github.com/inneroot/telenotify/pkg/api/grpc"
	"google.golang.org/grpc"
)

type grpcServerAPI struct {
	telebot_notify.UnimplementedNotifyServiceServer
	ns *notify_service.NotifyService
}

func New(ns *notify_service.NotifyService) *grpcServerAPI {
	return &grpcServerAPI{
		ns: ns,
	}
}

func (s *grpcServerAPI) Notify(ctx context.Context, req *telebot_notify.NotifyRequest) (*telebot_notify.NotifyResponse, error) {
	slog.Info("notify", slog.String("message", req.Message))
	err := s.ns.Notify(ctx, req.Message)
	return &telebot_notify.NotifyResponse{}, err
}

func RegisterNotifyServiceServer(gRPC *grpc.Server, api *grpcServerAPI) {
	telebot_notify.RegisterNotifyServiceServer(gRPC, api)
}
