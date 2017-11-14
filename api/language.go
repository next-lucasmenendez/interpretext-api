package api

import "github.com/abadojack/whatlanggo"

func getLanguage(input string) (lang string, ok bool) {
	info := whatlanggo.Detect(input)
	l := whatlanggo.LangToString(info.Lang)
	if l == "spa" {
		return "es", true
	} else if l == "eng" {
		return "en", true
	}

	return "", false
}