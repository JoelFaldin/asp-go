package main

import (
	"asp-go/internal/config"
	"asp-go/internal/crawler"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Inicializar logger:
	logHandler := logger.New(false)

	// Inicializar godotenv:
	err := godotenv.Load()
	if err != nil {
		logHandler.Fatal(err)
	}

	// Inicializar la db:
	dbURL := config.GetDBURL()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logHandler.Fatal(err)
	}
	store := storage.New(pool)
	defer pool.Close()

	// test:
	i, err := store.Save(ctx, "https://google.com/")
	if err != nil {
		logHandler.Fatal(err)
	}
	fmt.Println(i)

	// Cargar el archivo:
	sites, err := config.GetFile()
	if err != nil {
		logHandler.Fatal(err)
	}

	startOp := time.Now()

	// Ejecutar lógica principal:
	crawler.Crawl(ctx, logHandler, sites)

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)
}
