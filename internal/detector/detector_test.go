package detector

import (
	"asp-go/internal/fetcher"
	"net/http"
	"testing"
)

func TestDetectByServerHeader(t *testing.T) {
	res := fetcher.Result{}

	headers := http.Header{}
	headers.Set("Server", "cloudflare")

	res.Headers = headers

	r := detectUsingServerHeader(&res)

	if len(r) != 1 {
		t.Errorf("Expected 1 detection, got: %d", len(r))
	}
	if r[0].Name != "cloudflare" {
		t.Errorf("Expected detection name = cloudflare, got: %s", r[0].Name)
	}
}
