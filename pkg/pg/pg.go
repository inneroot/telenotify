package pg

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	var (
		err    error
		pgOnce sync.Once
		pool   *pgxpool.Pool
	)

	pgOnce.Do(func() {
		pool, err = pgxpool.New(ctx, connString)
		if err != nil {
			err = fmt.Errorf("unable to create connection pool: %w", err)
		}
	})

	err = WaitConnection(ctx, pool, 30, time.Second)

	return pool, err
}

func Ping(ctx context.Context, pool *pgxpool.Pool) error {
	return pool.Ping(ctx)
}

func ClosePool(pool *pgxpool.Pool) {
	pool.Close()
}

func WaitConnection(ctx context.Context, pool *pgxpool.Pool, retryLimit int, retryTimeout time.Duration) error {
	for i := range retryLimit {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			pingErr := Ping(ctx, pool)
			if pingErr != nil {
				if i >= retryLimit-1 {
					return pingErr
				} else {
					slog.Info(fmt.Sprintf("db did not responded. retry in %s...", retryTimeout))
					time.Sleep(retryTimeout)
				}
			}
		}
	}

	return nil
}
