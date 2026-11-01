package main

import (
	"asp-go/internal/config"
	"asp-go/internal/crawler"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Parsear flags:
	flags := config.ParseFlags()

	// Inicializar logger:
	logHandler := logger.New(false)

	if err := run(flags, logHandler); err != nil {
		logHandler.Fatal(err)
	}
}

func run(flags config.Flags, logHandler *logger.Logger) error {
	// Inicializar un contexto para las operaciones:
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Inicializar godotenv:
	err := godotenv.Load()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("configuring env: %w", err)
	}

	// Inicializar la db:
	dbURL := config.GetDBURL(flags.DB)
	pool, err := pgxpool.New(ctx, dbURL)
	error := pool.Ping(ctx)
	if err != nil || error != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	store := storage.New(pool)
	defer pool.Close()

	// Cargar el archivo:
	sites, err := config.GetFile(flags.SitesPath)
	if err != nil {
		return fmt.Errorf("loading domains file: %w", err)
	}

	startOp := time.Now()

	// Ejecutar lógica principal:
	crawler.Crawl(ctx, flags.DryRun, flags.Workers, flags.Timeout, logHandler, store, sites)

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)

	return nil
}
