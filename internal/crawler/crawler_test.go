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
		t.Errorf("expected between 9ms and 1.5s, got %d", elapsed)
	}
}
