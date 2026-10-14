package detector

import (
	"asp-go/internal/fetcher"
	"strings"
)

// Detecta patrones comunes de frameworks/librerias de frontend
func detectByBodyPatterns(res *fetcher.Result) []Detection {
	dt := []Detection{}

	for _, r := range bPatterns {
		bodyContains := strings.Contains(r.PatternString, res.Body)

		if bodyContains {
			dt = append(dt, Detection{Category: r.Category, TechName: r.Name})
		}
	}

	return dt
}

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
