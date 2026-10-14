package detector

import "regexp"

type Detection struct {
	Category string
	TechName string
}

var rgx = regexp.MustCompile(`(?i)<meta\s+name=(?:"|')generator(?:"|')\s+(content=(?:"|')([^"']+)(?:"|'))\s*/>`)

type HeaderPatterns struct {
	HeaderContains string
	Category       string
	Name           string
}

var hPatterns = []HeaderPatterns{
	{HeaderContains: "nginx", Category: "web-server", Name: "nginx"},
	{HeaderContains: "apache", Category: "web-server", Name: "apache"},
	{HeaderContains: "cloudflare", Category: "cdn", Name: "cloudflare"},
}

type BodyPatterns struct {
	PatternString string
	PatternRegex  *regexp.Regexp
	Category      string
	Name          string
}

var bPatterns = []BodyPatterns{
	{PatternString: "__NEXT_DATA__", Category: "framework", Name: "Next.js"},
	{PatternString: "/_next/static/", Category: "framework", Name: "Next.js"},
	{PatternString: "data-reactroot", Category: "ui library", Name: "React"},
	{PatternString: "/_nuxt/", Category: "framework", Name: "Nuxt"},
	{PatternString: "data-v", Category: "framework", Name: "Vue"},
}
