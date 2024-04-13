package repository

import "context"

type SubscriberRepository interface {
	GetAll(ctx context.Context) ([]int, error)
	Add(ctx context.Context, id int) error
	Del(ctx context.Context, id int) error
	Close()
}
