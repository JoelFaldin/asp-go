package detector

import (
	"asp-go/internal/fetcher"
)

// Tipo de firma específico para las funciones:
type rule func(*fetcher.Result) []Detection

var rules = []rule{
	detectUsingServerHeader,
	detectByPoweredByHeader,
	detectMetaTag,
	detectByBodyPatterns,
}

// Función principal del archivo.
// Dado un conjunto de Headers, ejecuta cada función de este archivo y devuelve el resultado.
func Detect(res *fetcher.Result) PageInfo {
	found := []Detection{}

	for _, r := range rules {
		det := r(res)

		found = append(found, det...)
	}

	title := DetectSiteTitle(res)
	pageInfo := PageInfo{Title: title, Detections: found}

	return pageInfo
}
