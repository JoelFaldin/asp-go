package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"context"
	"fmt"
	"time"
)

// Procesa un batch de numWorkers urls
func worker(id int, jobs <-chan config.Site, summary *Summary, limiter *Limiter, logHandler *logger.Logger) {
	for site := range jobs {
		startReq := time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		r, err := fetcher.Fetch(ctx, site.URL)
		duration := time.Since(startReq)

		if err != nil {
			logHandler.Error(err)
			summary.IncrementFailures()
			continue
		}

		res := detector.Detect(r)
		msg := fmt.Sprintf("%d %s", id, res)
		logHandler.SuccessTime(msg, duration)

		summary.IncrementSuccess()
	}
}
