package detector

import (
	"asp-go/internal/fetcher"
	"net/http"
	"testing"
)

// Structs y slice de cases con distintos casos para probar la funcion detectUsingServerHeader:
type testCase struct {
	name           string
	serverHeader   string
	wantCategory   string
	wantName       string
	wantDetections int
}

var headerCases = []testCase{
	{name: "cloudflare detected", serverHeader: "cloudflare", wantCategory: "cdn", wantName: "cloudflare", wantDetections: 1},
	{name: "nginx detected", serverHeader: "nginx/1.18.0", wantCategory: "web-server", wantName: "nginx", wantDetections: 1},
	{name: "no header doesnt break", serverHeader: "", wantDetections: 0},
	{name: "unknown header doesnt match", serverHeader: "localServer:/1.0", wantDetections: 0},
}

// Ejecuta la funcion detectUsingServerHeader por cada case del slice:
func TestDetectByServerHeader(t *testing.T) {
	for _, tc := range headerCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fetcher.Result{}

			headers := http.Header{}
			headers.Set("Server", tc.serverHeader)

			res.Headers = headers

			r := detectUsingServerHeader(&res)

			if len(r) != tc.wantDetections {
				t.Errorf("expected %d detections, got %d", tc.wantDetections, len(r))
			}
			if len(r) < 1 {
				return
			}
			if r[0].Category != tc.wantCategory {
				t.Errorf("expected category %s, got %s", tc.wantCategory, r[0].Category)
			}
			if r[0].Name != tc.wantName {
				t.Errorf("expected name %s, got %s", tc.wantName, r[0].Name)
			}
		})
	}

}

var poweredByCases = []testCase{
	{name: "php detected", serverHeader: "php", wantCategory: "backend", wantName: "php", wantDetections: 1},
	{name: "next.js detected", serverHeader: "next.js", wantCategory: "frontend", wantName: "next.js", wantDetections: 1},
	{name: "no header doesnt break", serverHeader: "", wantDetections: 0},
	{name: "unknown header doenst match", serverHeader: "localServer/1.1", wantDetections: 0},
}

func TestDetectByPoweredByHeader(t *testing.T) {
	for _, tc := range poweredByCases {
		res := fetcher.Result{}

		headers := http.Header{}
		headers.Set("X-Powered-By", tc.serverHeader)

		res.Headers = headers

		r := detectByPoweredByHeader(&res)

		if len(r) != tc.wantDetections {
			t.Errorf("expected %d detections, got %d", tc.wantDetections, len(r))
		}
		if len(r) < 1 {
			return
		}
		if r[0].Category != tc.wantCategory {
			t.Errorf("expected %s category, got %s", tc.wantCategory, r[0].Category)
		}
		if r[0].Name != tc.wantName {
			t.Errorf("exptected %s name, got %s", tc.wantName, r[0].Category)
		}
	}
}
