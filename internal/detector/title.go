package detector

import (
	"asp-go/internal/fetcher"
	"html"
	"regexp"
)

var titleRgx = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)

// Obtiene el title del sitio
func DetectSiteTitle(res *fetcher.Result) string {
	body := res.Body

	matches := titleRgx.FindStringSubmatch(body)
	if len(matches) < 2 {
		return ""
	}

	unescapedTitle := html.UnescapeString(matches[1])

	return unescapedTitle
}
