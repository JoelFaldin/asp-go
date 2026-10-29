package storage

import "asp-go/internal/detector"

type Scan struct {
	SiteID     int
	StatusCode *int
	Title      *string
	Detections []detector.Detection
	HostName   string
}
