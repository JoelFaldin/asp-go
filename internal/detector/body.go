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
// Extrae el contenido del tag y lo devuelve en un slice.
func detectMetaTag(res *fetcher.Result) []Detection {
	body := res.Body

	dt := []Detection{}
	matches := metaRgx.FindAllStringSubmatch(body, -1)
	if matches == nil {
		return dt
	}

	for _, item := range matches {
		m := metaGeneratorRgx.FindStringSubmatch(item[0])

		if len(m) > 1 {
			for _, r := range bTagPatterns {
				bodyContains := strings.Contains(m[1], r.PatternString)
				version := pluginVerRgx.FindString(m[1])

				if bodyContains {
					dt = append(dt, Detection{Category: r.Category, TechName: r.Name, Version: version})
				}
			}
		}
	}

	return dt
}
