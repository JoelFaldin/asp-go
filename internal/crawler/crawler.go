package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/logger"
	"fmt"
	"sync"
)

// Se encarga de la lógica principal (manda requests a los sitios y procesa la respuesta)
// Utiliza goroutines y channels para limitar un número concreto de goroutines que se pueden lanzar a la vez
func Crawl(logHandler *logger.Logger, sites []config.Site) {
	summary := NewSummary()
	limiters := NewLimiters()

	var wg sync.WaitGroup

	jobs := make(chan config.Site, len(sites))

	numWorkers := 5
	// Lanzar los numWorkers goroutines y se quedan escuchando el channel jobs
	for i := range numWorkers {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, jobs, summary, limiters, logHandler)
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
