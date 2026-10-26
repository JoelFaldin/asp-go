package config

import (
	"encoding/json"
	"log"
	"os"
)

type Site struct {
	URL string `json:"url"`
}

func GetFile() ([]Site, error) {
	data, err := os.ReadFile("./data/domains.json")
	if err != nil {
		return nil, err
	}

	var sites []Site

	if err := json.Unmarshal(data, &sites); err != nil {
		return nil, err
	}

	return sites, nil
}

func GetDBURL() string {
	url := os.Getenv("DATABASE_URL")
	if len(url) == 0 {
		log.Fatalf("Database url env variable not found")
	}

	return url
}
