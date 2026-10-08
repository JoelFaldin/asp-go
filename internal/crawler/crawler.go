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

// Se encarga de la lógica principal (manda requests a los sitios y procesa la respuesta)
// Utiliza goroutines y channels para limitar un número concreto de goroutines que se pueden lanzar a la vez
func Crawl(logHandler *logger.Logger, sites []config.Site) {
	summary := Summary{}

	var wg sync.WaitGroup

	jobs := make(chan config.Site, len(sites))

	numWorkers := 5
	// Lanzar los numWorkers goroutines y se quedan escuchando el channel jobs
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, jobs, &summary, logHandler)
		}(i)
	}

	// Llenar el channel jobs con urls
	for _, n := range sites {
		jobs <- n
	}

	close(jobs)
	wg.Wait()

	// Imprimir resumen:
	msg1 := fmt.Sprintf("Successes: %d", summary.successful)
	logHandler.Info(msg1)
	msg2 := fmt.Sprintf("Failures: %d", summary.failures)
	logHandler.Info(msg2)
}

// Procesa un batch de numWorkers urls
func worker(id int, jobs <-chan config.Site, summary *Summary, logHandler *logger.Logger) {
	for site := range jobs {
		startReq := time.Now()
		r, err := fetcher.Fetch(site.URL)
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
