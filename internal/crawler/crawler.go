package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"fmt"
	"sync"
	"time"
)

type Saver interface {
	SaveScan(ctx context.Context, sc storage.Scan) error
}

var _ Saver = (*storage.Store)(nil)

// Se encarga de la lógica principal (manda requests a los sitios y procesa la respuesta)
// Utiliza goroutines y channels para limitar un número concreto de goroutines que se pueden lanzar a la vez:
func Crawl(ctx context.Context, workers int, timeout time.Duration, logHandler *logger.Logger, store Saver, sites []config.Site) {
	summary := NewSummary()
	limiters := NewLimiters()

	var wg sync.WaitGroup

	jobs := make(chan config.Site, len(sites))
	results := make(chan storage.Scan, 10)
	done := make(chan struct{})

	// Lanza goroutines que escriben a la base de datos:
	go func() {
		defer close(done)
		for sc := range results {
			saveCtx := context.WithoutCancel(ctx)
			timeoutCtx, cancel := context.WithTimeout(saveCtx, 5*time.Second)

			err := store.SaveScan(timeoutCtx, sc)
			defer cancel()
			if err != nil {
				logHandler.Error(err)
				continue
			}
		}
	}()

	// Lanzar los numWorkers goroutines y se quedan escuchando el channel jobs
	for i := range workers {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(ctx, timeout, id, jobs, results, summary, limiters, logHandler)
		}(i)
	}

	// Llenar el channel jobs con urls
	for _, n := range sites {
		jobs <- n
	}

	close(jobs)
	wg.Wait()
	close(results)
	<-done

	// Imprimir resumen:
	msg1 := fmt.Sprintf("Successes: %d", summary.successful)
	logHandler.Info(msg1)
	msg2 := fmt.Sprintf("Failures: %d", summary.failures)
	logHandler.Info(msg2)
}
