package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"
)

// Procesa un batch de numWorkers urls
func worker(ctx context.Context, id int, jobs <-chan config.Site, results chan<- storage.Scan, summary *Summary, limiter *Limiter, logHandler *logger.Logger, store *storage.Store) {
	for site := range jobs {
		if ctx.Err() != nil {
			break
		}

		startReq := time.Now()

		err := limiter.Wait(ctx, site.URL)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return
			}

			logHandler.Error(err)
			summary.IncrementFailures()
			continue
		}

		ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)

		r, err := fetcher.Fetch(ctxTimeout, site.URL)
		cancel()

		duration := time.Since(startReq)

		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return
			}

			logHandler.Error(err)
			summary.IncrementFailures()
			continue
		}

		res := detector.Detect(r)

		// Guardar resultados en la base de datos:
		parsedURL, err := url.Parse(r.URL)
		if err != nil {
			logHandler.Error(err)
			return
		}
		scan := storage.Scan{
			StatusCode: &r.StatusCode,
			Title:      &res.Title,
			Detections: res.Detections,
			HostName:   parsedURL.String(),
		}

		results <- scan

		// Imprime resultados en la consola:
		msg := fmt.Sprintf("%d %s", id, res)
		logHandler.SuccessTime(msg, duration)

		summary.IncrementSuccess()
	}
}
