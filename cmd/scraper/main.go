package main

import (
	"asp-go/internal/config"
	"asp-go/internal/crawler"
	"asp-go/internal/logger"
	"asp-go/internal/storage"
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Inicializar un contexto para las operaciones:
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Configurar flags:
	dryRun := flag.Bool("dry-run", false, "prints results to the console and doesnt save data into the database")
	workers := flag.Int("workers", 5, "amount of concurrent workers")
	timeout := flag.Duration("timeout", 10*time.Second, "timeout per request")
	sitesPath := flag.String("sites", "./data/domains.json", "route to domains file")
	flag.Parse()

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
	sites, err := config.GetFile(*sitesPath)
	if err != nil {
		logHandler.Fatal(err)
	}

	startOp := time.Now()

	// Ejecutar lógica principal:
	crawler.Crawl(ctx, *dryRun, *workers, *timeout, logHandler, store, sites)

	totalTime := time.Since(startOp)
	logHandler.Duration("Total operation time:", totalTime)
}
