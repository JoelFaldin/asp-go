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
	Body       string
}

func Fetch(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{
		URL:        res.Request.URL.Host,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       string(body),
	}

	return r, nil
}
