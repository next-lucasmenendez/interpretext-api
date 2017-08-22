package api

import (
	a "github.com/lucasmenendez/gobstract"

)

func getSummary(input, lang string) map[string]interface{} {
	if abstract, err := a.NewAbstract(input, lang); err != nil {
		return nil
	} else {
		var highlights []string = abstract.GetHightlights()
		var keywords []string = abstract.GetKeywords()

		return map[string]interface{} {
			"keywords": keywords,
			"highlights": highlights,
		}
	}
	return nil
}