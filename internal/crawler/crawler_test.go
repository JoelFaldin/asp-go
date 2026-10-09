package crawler

import (
	"context"
	"testing"
	"time"
)

func TestLimiterSameDomain(t *testing.T) {
	l := NewLimiters()
	ctx := context.Background()

	start := time.Now()

	l.Wait(ctx, "https://www.a.cl/")
	l.Wait(ctx, "https://www.a.cl/content")

	elapsed := time.Since(start)

	if elapsed < 900*time.Millisecond || elapsed > 1500*time.Millisecond {
		t.Errorf("expected between 9ms and 1.5s, got %v", elapsed)
	}
}

func TestLimiterDifferentDomains(t *testing.T) {
	l := NewLimiters()
	ctx := context.Background()

	start := time.Now()

	l.Wait(ctx, "https://www.a.cl/")
	l.Wait(ctx, "https://www.b.cl/")
	l.Wait(ctx, "https://www.c.cl/")

	elapsed := time.Since(start)

	if elapsed > 500*time.Millisecond {
		t.Errorf("expected less than 500ms, got %v", elapsed)
	}
}
