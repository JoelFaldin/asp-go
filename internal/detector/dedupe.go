package detector

import (
	"strings"
)

// Normaliza el resultado de un []Detection.
// El []Detection entrante puede tener resultados duplicados, esta función los normaliza.
func Dedupe(detect []Detection) []Detection {
	var result []Detection
	seen := make(map[string]int)

	for _, d := range detect {
		name := strings.ToLower(d.TechName)

		// Revisa si seen[name] ya existe en el map.
		// Si ya existe, agrega el source al slice existente:
		if idx, ok := seen[name]; ok {
			result[idx].Sources = append(result[idx].Sources, d.Sources...)

			continue
		}

		result = append(result, d)
		seen[name] = len(result) - 1
	}

	return result
}
