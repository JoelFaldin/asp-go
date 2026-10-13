package detector

import (
	"asp-go/internal/fetcher"
	"fmt"
	"html"
	"regexp"
	"strings"
)

type Detection struct {
	Category string
	TechName string
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

var rgx = regexp.MustCompile(`(?i)<meta\s+name=(?:"|')generator(?:"|')\s+(content=(?:"|')([^"']+)(?:"|'))\s*/>`)

// Usa el regex para matchear la tag <meta>, muy común en sitios cms.
func detectMetaTag(res *fetcher.Result) []Detection {
	body := res.Body

	matches := rgx.FindStringSubmatch(body)

	dt := []Detection{}
	if matches == nil {
		return dt
	}

	dt = append(dt, Detection{Category: "cms", TechName: matches[2]})
	return dt
}

var rgxBdyPttrn = regexp.MustCompile(``)

type BodyPatterns struct {
	PatternString string
	PatternRegex  *regexp.Regexp
	Category      string
	Name          string
}

var patterns = []BodyPatterns{
	{PatternString: "__NEXT_DATA__", Category: "framework", Name: "Next.js"},
	{PatternString: "/_next/static/", Category: "framework", Name: "Next.js"},
	{PatternString: "data-reactroot", Category: "ui library", Name: "React"},
	{PatternString: "/_nuxt/", Category: "framework", Name: "Nuxt"},
	{PatternString: "data-v", Category: "framework", Name: "Vue"},
}

// Detecta patrones comunes de frameworks/librerias de frontend
func detectByBodyPatterns(res *fetcher.Result) []Detection {
	dt := []Detection{}

	for _, r := range patterns {
		bodyContains := strings.Contains(r.PatternString, res.Body)

		if bodyContains {
			dt = append(dt, Detection{Category: r.Category, TechName: r.Name})
		}
	}

	return dt
}

var titleRgx = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)

// Obtiene el title del sitio
func detectSiteTitle(res *fetcher.Result) string {
	body := res.Body

	matches := titleRgx.FindStringSubmatch(body)
	unescapedTitle := html.UnescapeString(matches[1])

	return unescapedTitle
}

var rules = []rule{
	detectUsingServerHeader,
	detectByPoweredByHeader,
	detectMetaTag,
	detectByBodyPatterns,
}

// Función principal del archivo.
// Dado un conjunto de Headers, ejecuta cada función de este archivo y devuelve el resultado.
func Detect(res *fetcher.Result) []Detection {
	found := []Detection{}

	for _, r := range rules {
		det := r(res)

		found = append(found, det...)
	}

	title := detectSiteTitle(res)
	fmt.Println("Site title: ", title)

	return found
}
