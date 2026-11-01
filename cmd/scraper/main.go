package main

import (
	"asp-go/internal/config"
	"asp-go/internal/crawler"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	// Inicializar un contexto para las operaciones:
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	flags := config.ParseFlags()

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

	// Cargar el archivo:
	sites, err := config.GetFile(*flags.SitesPath)
	if err != nil {
		logHandler.Fatal(err)
	}

	startOp := time.Now()

	// Ejecutar lógica principal:
	crawler.Crawl(ctx, *flags.DryRun, *flags.Workers, *flags.Timeout, logHandler, store, sites)

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)

	return nil
}
