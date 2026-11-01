package config

import (
	"flag"
	"time"
)

type Flags struct {
	DryRun    bool
	Workers   int
	Timeout   time.Duration
	SitesPath string
	DB        string
}

func ParseFlags() Flags {
	// Configurar flags:
	dryRun := flag.Bool("dry-run", false, "prints results to the console and doesnt save data into the database")
	workers := flag.Int("workers", 5, "amount of concurrent workers")
	timeout := flag.Duration("timeout", 10*time.Second, "timeout per request")
	sitesPath := flag.String("sites", "./data/domains.json", "route to domains file")
	db := flag.String("db", "", "connection string for the database")
	flag.Parse()

	flags := Flags{
		DryRun:    *dryRun,
		Workers:   *workers,
		Timeout:   *timeout,
		SitesPath: *sitesPath,
		DB:        *db,
	}

	return flags
}
