package main

import (
	"asp-go/internal/detector"
	"asp-go/internal/fetcher"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	for i := 1; i < len(args); i++ {
		r, err := fetcher.Fetch(args[i])
		if err != nil {
			log.Println(err)
			return
		}

		d := detector.DetectUsingServerHeader(r)
		fmt.Println(d)
	}
}
