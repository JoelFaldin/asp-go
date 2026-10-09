package main

import (
	"asp-go/internal/config"
	"asp-go/internal/crawler"
	"asp-go/internal/logger"
	"context"
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

	ctx := context.Background()
	// Ejecutar lógica principal:
	crawler.Crawl(ctx, logHandler, sites)

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)
}
