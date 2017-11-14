package gocountrydetect

import (
	"regexp"
	"strings"
)

type Country struct {
	Code, Name string
}

type Countries []Country

func (countries Countries) contains(country Country) bool {
	for _, c := range countries {
		if c.Name == country.Name && c.Code == country.Code {
			return true
		}
	}

	return false
}

func Detect(text, lang string) Countries {
	var countries Countries = getCountries(lang)
	var words []string = tokenize(text)

	var results Countries
	for _, w := range words {
		for _, c := range countries {
			if w == strings.ToLower(c.Name) {
				if !results.contains(c) {
					results = append(results, c)
				}
			}
		}
	}

	return results
}

func tokenize(text string) []string {
	var words []string

	var rgxClean *regexp.Regexp= regexp.MustCompile(`\[|]|\(|\)|\{|}|“|”|«|»|,|´|’|-|_|—|\.\.|:`)
	var cleaned string = rgxClean.ReplaceAllString(text, "")

	var rgx_word = regexp.MustCompile(`\s`)
	var parts []string = rgx_word.Split(cleaned, -1)

	for _, raw_word := range parts {
		var word string = strings.TrimSpace(raw_word)
		if len(word) > 3 {
			words = append(words, strings.ToLower(word))
		}
	}
	return words
}
