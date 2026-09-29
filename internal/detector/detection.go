package detector

import (
	"asp-go/internal/fetcher"
	"strings"
)

type Detection struct {
	Category string
	Name     string
}

func DetectUsingServerHeader(res *fetcher.Result) []Detection {
	sv := res.Headers.Get("Server")

	if len(sv) == 0 {
		return nil
	}

	lower := strings.ToLower(sv)

	dt := []Detection{}
	if strings.Contains(lower, "nginx") {
		dt = append(dt, Detection{Category: "web-server", Name: "nginx"})
	}
	if strings.Contains(lower, "apache") {
		dt = append(dt, Detection{Category: "web-server", Name: "apache"})
	}
	if strings.Contains(lower, "cloudflare") {
		dt = append(dt, Detection{Category: "cdn", Name: "cloudflare"})
	}

	return dt
}
