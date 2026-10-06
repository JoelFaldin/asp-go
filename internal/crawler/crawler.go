package crawler

import (
	"asp-go/internal/config"
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"asp-go/internal/logger"
	"fmt"
	"time"
)

func Crawl(logHandler *logger.Logger, sites []config.Site) {
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
		logHandler.SuccessTime(msg, duration)
	}
}
