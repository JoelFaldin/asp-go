package detector

import (
	"asp-go/internal/fetcher"
	"strings"
)

type Detection struct {
	Category string
	Name     string
}

type rule func(*fetcher.Result) []Detection

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

func DetectByPoweredByHeader(res *fetcher.Result) []Detection {
	poweredBy := res.Headers.Get("X-Powered-By")

	if len(poweredBy) == 0 {
		return nil
	}

	lower := strings.ToLower(poweredBy)

	dt := []Detection{}
	if strings.Contains(lower, "php") {
		dt = append(dt, Detection{Category: "stack", Name: "php"})
	}
	if strings.Contains(lower, "express") {
		dt = append(dt, Detection{Category: "stack", Name: "express"})
	}
	if strings.Contains(lower, "asp-net") {
		dt = append(dt, Detection{Category: "stack", Name: "asp-net"})
	}

	return dt
}

var rules = []rule{
	DetectUsingServerHeader,
	DetectByPoweredByHeader,
}
