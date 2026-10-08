package fetcher

import (
	"context"
	"io"
	"net/http"
)

type Result struct {
	URL        string
	StatusCode int
	Headers    http.Header
	Cookies    []string
	Body       string
}

// Hace la request a la web.
// Almacena lo encontrado en el struct Result
func Fetch(ctx context.Context, url string) (*Result, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Result{
		URL:        res.Request.URL.Host,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       string(body),
	}

	return r, nil
}
