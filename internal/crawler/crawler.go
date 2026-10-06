package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"fmt"
	"sync"
	"time"
)

type Summary struct {
	successful int
	failures   int
	mu         sync.Mutex
}

func Crawl(logHandler *logger.Logger, sites []config.Site) {
	summary := Summary{}

	for _, s := range sites {
		startReq := time.Now()
		r, err := fetcher.Fetch(s.URL)
		duration := time.Since(startReq)

		if err != nil {
			logHandler.Error(err)
			summary.IncrementFailures()
			continue
		}

		res := detector.Detect(r)
		msg := fmt.Sprintf("%s", res)
		logHandler.SuccessTime(msg, duration)

		summary.IncrementSuccess()
	}

	// Imprimir resumen:
	fmt.Printf("success: %d, failures: %d\n", summary.successful, summary.failures)
}

func (s *Summary) IncrementSuccess() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.successful++
}

func (s *Summary) IncrementFailures() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.failures++
}
