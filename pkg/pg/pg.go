package pg

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	retryLimit   = 10
	retryTimeout = 5 * time.Second
)

type postgres struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, logger *slog.Logger, connString string) (*postgres, error) {
	log := logger.With(slog.String("module", "pg"))
	var dbErr error

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			dbErr = fmt.Errorf("unable to create connection pool: %w", err)
		}

		pgInstance = &postgres{db, log}
	})

	return pgInstance, dbErr
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}

func (pg *postgres) WaitConnection(ctx context.Context) error {
	for i := range retryLimit {
		pingErr := pg.Ping(ctx)
		if pingErr != nil {
			if i >= retryLimit-1 {
				return pingErr
			} else {
				pg.log.Info(fmt.Sprintf("db did not responded. retry in %s...", retryTimeout))
				time.Sleep(retryTimeout)
			}
		}
	}

	return nil
}
