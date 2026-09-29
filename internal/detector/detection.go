package detector

import (
	"asp-go/internal/fetcher"
	"fmt"
	"strings"
)

type Detection struct {
	Framework string
	Cms       string
	Backend   string
}

func DetectUsingServerHeader(res *fetcher.Result) {
	sv := res.Headers.Get("Server")

	if len(sv) == 0 {
		return
	}

	lower := strings.ToLower(sv)

	if strings.Contains(lower, "nginx") {
		fmt.Println("is using nginx")
	} else if strings.Contains(lower, "apache") {
		fmt.Println("is using apache")
	} else if strings.Contains(lower, "cloudflare") {
		fmt.Println("is using cloudflare")
	}
}
