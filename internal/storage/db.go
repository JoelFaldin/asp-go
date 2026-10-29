package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) Save(ctx context.Context, domain string) (int, error) {
	query := `INSERT INTO sites (domain) VALUES ($1)`
	res, err := s.pool.Exec(ctx, query, domain)
	if err != nil {
		return 0, err
	}

	fmt.Println(res)
	return 1, nil
}
