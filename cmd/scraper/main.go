package main

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"fmt"
	"log"
	"time"
)

func main() {
	// Get data file:
	sites, err := config.GetFile()
	if err != nil {
		log.Fatalf("couldnt load file:", err)
	}

	// Initialize logger:
	logHandler := logger.New(false)

	for _, s := range sites {
		startReq := time.Now()
		r, err := fetcher.Fetch(s.URL)
		duration := time.Since(startReq)

		if err != nil {
			log.Println("there was an error loading a site:", err)
			continue
		}

		res := detector.Detect(r)
		msg := fmt.Sprintf("%s", res)
		logHandler.Info(msg, duration)
	}
}
