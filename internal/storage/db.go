package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	pool.Ping(ctx)

	return pool, err
}
