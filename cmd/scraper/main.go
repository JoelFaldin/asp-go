package main

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"fmt"
	"time"
)

func main() {
	// Inicializar logger:
	logHandler := logger.New(false)

	// Cargar el archivo:
	sites, err := config.GetFile()
	if err != nil {
		logHandler.Fatal(err)
	}

	startOp := time.Now()

	for _, s := range sites {
		startReq := time.Now()
		r, err := fetcher.Fetch(s.URL)
		duration := time.Since(startReq)

		if err != nil {
			logHandler.Error(err)
			continue
		}

		res := detector.Detect(r)
		msg := fmt.Sprintf("%s", res)
		logHandler.Success(msg, duration)
	}

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)
}
