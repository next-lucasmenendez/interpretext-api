package api

import "github.com/chrisport/go-lang-detector/langdet/langdetdef"

func getLanguage(input string) (lang string, ok bool) {
	detector := langdetdef.NewWithDefaultLanguages()
	l := detector.GetClosestLanguage(input)

	if l == "spanish" {
		lang = "es"
	} else if l == "english" {
		return "", false
	}

	return lang, true
}