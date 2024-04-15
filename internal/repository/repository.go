package repository

import "context"

type SubscriberRepository interface {
	GetAll(ctx context.Context) ([]int64, error)
	Add(ctx context.Context, id int64) error
	Del(ctx context.Context, id int64) error
	Close()
}
