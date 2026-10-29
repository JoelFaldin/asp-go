package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) Save(ctx context.Context, domain string) (int, error) {
	query := `
		INSERT INTO sites (domain, last_scanned_at)
		VALUES ($1, now())
		ON CONFLICT (domain)
		DO UPDATE SET
			domain = EXCLUDED.domain,
			last_scanned_at = now()
		RETURNING id
	`
	var id int
	err := s.pool.QueryRow(ctx, query, domain).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
