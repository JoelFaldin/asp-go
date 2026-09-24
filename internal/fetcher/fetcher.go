package fetcher

import (
	"io"
	"net/http"
)

type Response struct {
	URL        string
	StatusCode int
	Headers    http.Header
	Cookies    []string
	Boydy      string
}

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
