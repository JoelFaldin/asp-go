package storage

import "asp-go/internal/detector"

// Struct utilizado para insertar datos en la db mediante la función SaveScan, en /internal/storage/db.go:
type Scan struct {
	SiteID     int
	StatusCode *int
	Title      *string
	Detections []detector.Detection
	HostName   string
}
