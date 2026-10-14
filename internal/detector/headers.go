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

	for _, item := range hPatterns {
		if strings.Contains(lower, item.HeaderContains) {
			dt = append(dt, Detection{Category: item.Category, TechName: item.Name})
		}
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

	for _, item := range hPoweredByPatterns {
		if strings.Contains(lower, item.HeaderContains) {
			dt = append(dt, Detection{Category: item.Category, TechName: item.Name})
		}
	}

	return dt
}
