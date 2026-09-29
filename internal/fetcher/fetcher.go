package fetcher

import (
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
func Fetch(url string) (*Result, error) {
	res, err := http.Get(url)
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
