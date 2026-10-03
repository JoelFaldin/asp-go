package main

import (
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Site struct {
	URL string `json:"url"`
}

func main() {
	data, err := os.ReadFile("./data/domains.json")
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}

	var sites []Site

	if err := json.Unmarshal(data, &sites); err != nil {
		fmt.Println("error parsing json", err)
		return
	}

	for _, s := range sites {
		r, err := fetcher.Fetch(s.URL)
		if err != nil {
			log.Println(err)
			return
		}

		det := detector.Detect(r)
		fmt.Println(det)
	}
}
