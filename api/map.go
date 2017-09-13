package api

import (
	"fmt"
	"strings"
	gcd "github.com/beeva-lucasmenendez/gocountrydetect"
)

const API string = "http://country-geojson-api.herokuapp.com"

func getMap(text, lang string) string {
	var countries gcd.Countries = gcd.Detect(text, lang)
	var names []string
	for _, c := range countries {
		names = append(names, strings.ToLower(c.Name))
	}

	if len(countries) == 0 {
		return ""
	}
	var plainNames string = strings.Join(names, ",")
	var endpoint string = fmt.Sprintf("%s/search/%s?lang=%s", API, plainNames, lang)
	return endpoint
}

