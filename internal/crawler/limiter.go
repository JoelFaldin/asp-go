package crawler

import (
	"context"
	"net/url"
	"sync"

	"golang.org/x/time/rate"
)

type Limiter struct {
	entry map[string]*rate.Limiter
	mu    sync.Mutex
}

func NewLimiters() *Limiter {
	return &Limiter{
		entry: make(map[string]*rate.Limiter),
	}
}

func (l *Limiter) Wait(ctx context.Context, rawURL string) error {
	l.mu.Lock()

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	host := parsedURL.Hostname()

	entry := l.getLimiter(host)
	return entry.Wait(ctx)
}

func (l *Limiter) getLimiter(host string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	val, ok := l.entry[host]
	if ok {
		return val
	}

	lim := rate.NewLimiter(1, 1)
	l.entry[host] = lim

	return lim
}
