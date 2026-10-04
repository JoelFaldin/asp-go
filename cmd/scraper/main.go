package main

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"fmt"
	"log"
)

func main() {
	// Get file:
	sites, err := config.GetFile()
	if err != nil {
		log.Fatalf("couldnt load file:", err)
	}

	for _, s := range sites {
		r, err := fetcher.Fetch(s.URL)
		if err != nil {
			log.Println("there was an error loading a site:", err)
			continue
		}

		det := detector.Detect(r)
		fmt.Println(det)
	}
}
