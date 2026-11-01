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

// Obtiene la URL de la database a la que se conectan:
func GetDBURL() string {
	url := os.Getenv("DATABASE_URL")
	if len(url) == 0 {
		log.Fatalf("Database url env variable not found")
	}

	return url
}
