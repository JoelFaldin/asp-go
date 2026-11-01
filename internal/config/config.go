package config

import (
	"encoding/json"
	"log"
	"os"
)

type Site struct {
	URL string `json:"url"`
}

// Obtiene el archivo de datos de donde saca los dominios:
func GetFile(sitesPath string) ([]Site, error) {
	data, err := os.ReadFile(sitesPath)
	if err != nil {
		return nil, err
	}

	var sites []Site

	if err := json.Unmarshal(data, &sites); err != nil {
		return nil, err
	}

	return sites, nil
}

// Obtiene la URL de la database a la que se conectan.
// Prioriza la dbURL que viene desde el flag. Si no está presente, lee la variable de entorno:
func GetDBURL(dbURL string) string {
	if len(dbURL) == 0 {
		url := os.Getenv("DATABASE_URL")
		if len(url) == 0 {
			log.Fatalf("Database url env variable not found")
		}

		return url
	} else {
		return dbURL
	}
}
