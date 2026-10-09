package crawler

import "sync"

// Gestiona el resumen de las operaciones.
// Las funciones IncrementSummary e IncrementFailures controlan su contenido
type Summary struct {
	successful int
	failures   int
	mu         sync.Mutex
}

func NewSummary() *Summary {
	return &Summary{
		successful: 0,
		failures:   0,
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
