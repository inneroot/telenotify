package notify_service

import "context"

type INotifier interface {
	NotifySubscribed(ctx context.Context, message string) error
}

type NotifyService struct {
	notifier INotifier
}

func New(notifier INotifier) *NotifyService {
	return &NotifyService{
		notifier,
	}
}

func (ns *NotifyService) Notify(ctx context.Context, message string) error {
	return ns.notifier.NotifySubscribed(ctx, message)
}
