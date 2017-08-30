package api

import (
	a "github.com/lucasmenendez/gobstract"

)

func getSummary(input, lang string, count int) map[string]interface{} {
	if abstract, err := a.NewAbstract(input, lang); err != nil {
		return nil
	} else {
		var highlights []string = abstract.GetHightlights(count)
		var bestSentence string = abstract.GetBestSentence()
		var keywords []string = abstract.GetKeywords()

		return map[string]interface{} {
			"best_sentence": bestSentence,
			"keywords": keywords,
			"highlights": highlights,
		}
	}
	return nil
}