package detector

import (
	"asp-go/internal/fetcher"
	"strings"
)

// Revisa el Header "Server" para detectar CDNs y servidores web comunes:
func detectUsingServerHeader(res *fetcher.Result) []Detection {
	sv := res.Headers.Get("Server")

	if len(sv) == 0 {
		return nil
	}

	lower := strings.ToLower(sv)

	dt := []Detection{}
	if strings.Contains(lower, "nginx") {
		dt = append(dt, Detection{Category: "web-server", TechName: "nginx"})
	}
	if strings.Contains(lower, "apache") {
		dt = append(dt, Detection{Category: "web-server", TechName: "apache"})
	}
	if strings.Contains(lower, "cloudflare") {
		dt = append(dt, Detection{Category: "cdn", TechName: "cloudflare"})
	}

	return dt
}

// Revisa el Header X-Powered-By para detectar servidores backends comunes:
func detectByPoweredByHeader(res *fetcher.Result) []Detection {
	poweredBy := res.Headers.Get("X-Powered-By")

	if len(poweredBy) == 0 {
		return nil
	}

	lower := strings.ToLower(poweredBy)

	dt := []Detection{}
	if strings.Contains(lower, "php") {
		dt = append(dt, Detection{Category: "backend", TechName: "php"})
	}
	if strings.Contains(lower, "express") {
		dt = append(dt, Detection{Category: "backend", TechName: "express"})
	}
	if strings.Contains(lower, "asp-net") {
		dt = append(dt, Detection{Category: "backend", TechName: "asp-net"})
	}
	if strings.Contains(lower, "next.js") {
		dt = append(dt, Detection{Category: "frontend", TechName: "next.js"})
	}

	return dt
}
