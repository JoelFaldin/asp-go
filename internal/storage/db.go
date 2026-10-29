package storage

import (
	"asp-go/internal/detector"
	"context"

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
		INSERT INTO scans (site_id, status_code, title, detections)
		VALUES ($1, $2, $3, $4)
	`

	// Revisar el slice detections:
	var dt []detector.Detection
	if sc.Detections == nil {
		dt = []detector.Detection{
			{Category: "framework", TechName: "Next.js", Sources: []string{"Header", "Body"}},
		}
	} else {
		dt = sc.Detections
	}

	_, err := s.pool.Exec(ctx, query, sc.SiteID, sc.StatusCode, sc.Title, dt)
	if err != nil {
		return err
	}

	return nil
}
