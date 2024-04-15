package pgrepo

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/inneroot/telenotify/internal/config"
	"github.com/inneroot/telenotify/pkg/pg"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Subscriber struct {
	RecipientID int `db:"recipient_id"`
}

type PGRepository struct {
	pool    *pgxpool.Pool
	log     *slog.Logger
	timeout time.Duration
}

func New(ctx context.Context, logger *slog.Logger, timeout time.Duration) (*PGRepository, error) {
	connStr := config.GetPGConnectionString()
	log := logger.With(slog.String("module", "PGRepository"))
	pool, err := pg.NewPGPool(ctx, connStr)
	return &PGRepository{pool, log, timeout}, err
}

func (pg *PGRepository) GetAll(ctx context.Context) ([]int64, error) {
	cctx, cancel := context.WithTimeout(ctx, pg.timeout)
	defer cancel()
	query := `SELECT recipient_id FROM recipients LIMIT 10000`

	var result []int64
	if err := pgxscan.Select(cctx, pg.pool, &result, query); err != nil {
		pg.log.Error("GetAll", "error", err.Error())
		return result, fmt.Errorf("failed to recipients ids, %w", err)
	}
	pg.log.Debug("GetAll", "result", result)
	return result, nil
}

func (pg *PGRepository) Add(ctx context.Context, id int64) error {
	cctx, cancel := context.WithTimeout(ctx, pg.timeout)
	defer cancel()

	query := `INSERT INTO recipients (recipient_id) VALUES ($1)`

	if _, err := pg.pool.Exec(cctx, query, id); err != nil {
		pg.log.Error("Add", "error", err.Error())
		return fmt.Errorf("failed to add recipient: %v", err)
	}

	pg.log.Debug("Add", "id", id)
	return nil
}

func (pg *PGRepository) Del(ctx context.Context, id int64) error {
	cctx, cancel := context.WithTimeout(ctx, pg.timeout)
	defer cancel()

	query := `DELETE FROM recipients WHERE recipient_id=($1)`

	if _, err := pg.pool.Exec(cctx, query, id); err != nil {
		pg.log.Error("Del", "error", err.Error())
		return fmt.Errorf("failed to del recipient: %v", err)
	}

	pg.log.Debug("Del", "id", id)
	return nil
}

func (pg *PGRepository) Close() {
	pg.pool.Close()
}
