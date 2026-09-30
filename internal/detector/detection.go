package detector

import (
	"asp-go/internal/fetcher"
	"regexp"
	"strings"
)

type Detection struct {
	Category string
	Name     string
}

// Tipo de firma específico para las funciones:
type rule func(*fetcher.Result) []Detection

// Revisa el Header "Server" para detectar CDNs y servidores web comunes:
func detectUsingServerHeader(res *fetcher.Result) []Detection {
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

// Revisa el Header X-Powered-By para detectar servidores backends comunes:
func detectByPoweredByHeader(res *fetcher.Result) []Detection {
	poweredBy := res.Headers.Get("X-Powered-By")

	if len(poweredBy) == 0 {
		return nil
	}

	lower := strings.ToLower(poweredBy)

	dt := []Detection{}
	if strings.Contains(lower, "php") {
		dt = append(dt, Detection{Category: "backend", Name: "php"})
	}
	if strings.Contains(lower, "express") {
		dt = append(dt, Detection{Category: "backend", Name: "express"})
	}
	if strings.Contains(lower, "asp-net") {
		dt = append(dt, Detection{Category: "backend", Name: "asp-net"})
	}

	return dt
}

var rgx = regexp.MustCompile(`(?i)<meta\s+name=(?:"|')generator(?:"|')\s+(content=(?:"|')([^"']+)(?:"|'))\s*/>`)

func detectMetaTag(res *fetcher.Result) []Detection {
	body := res.Body

	matches := rgx.FindStringSubmatch(body)

	dt := []Detection{}
	if matches == nil {
		return dt
	}

	dt = append(dt, Detection{Category: "cms", Name: matches[2]})
	return dt
}

var rules = []rule{
	detectUsingServerHeader,
	detectByPoweredByHeader,
	detectMetaTag,
}

// Función principal del archivo.
// Dado un conjunto de Headers, ejecuta cada función de este archivo y devuelve el resultado.
func Detect(res *fetcher.Result) []Detection {
	found := []Detection{}

	for _, r := range rules {
		det := r(res)

		found = append(found, det...)
	}

	return found
}
