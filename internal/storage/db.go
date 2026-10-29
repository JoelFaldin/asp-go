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

func (s *Store) UpsertSite(ctx context.Context, domain string) (int, error) {
	query := `
		INSERT INTO sites (domain, last_scanned_at)
		VALUES ($1, now())
		ON CONFLICT (domain)
		DO UPDATE SET
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

func (s *Store) InsertScan(ctx context.Context, sc Scan) error {
	query := `
		INSERT INTO scans (site_id, status_code, title)
		VALUES ($1, $2, $3)
	`
	res, err := s.pool.Exec(ctx, query, sc.SiteID, sc.StatusCode, sc.Title)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
