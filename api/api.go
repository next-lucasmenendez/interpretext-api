package api

import f "github.com/lucasmenendez/framework.go"

func MainHandler(c f.Context) {
	var ok bool
	var lang, input string
	if input, ok = c.FormValue("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
		return
	}

	if lang, ok = c.FormValue("lang"); !ok {
		if lang, ok = getLanguage(input); !ok {
			c.WriteErrorMessage("No correct language detected or provided.", 400)
			return
		}
	}

	if result := getSummary(input, lang); result != nil {
		c.JsonWrite(result, 200)
		return
	}

	c.WriteErrorMessage("Sorry! We didn't find any content :(", 404)
	return
}

func TweetHandler(c f.Context) {
	var ok bool
	var lang, input string
	if lang, ok = c.FormValue("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
		return
	}

	if input, ok = c.FormValue("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
		return
	}

	var tweet string = getTweet(input, lang)
	if tweet != "" {
		c.JsonWrite(map[string]string{"best_sentence": tweet}, 200)
	} else {
		c.WriteErrorMessage("Tweet not found :(", 404)
	}
	return
}

func SummaryHandler(c f.Context) {
	var ok bool
	var lang, input string
	if lang, ok = c.FormValue("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
	}

	if input, ok = c.FormValue("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
	}

	var data map[string]interface{} = getSummary(input, lang)
	if data != nil {
		c.JsonWrite(data, 200)
	} else {
		c.WriteErrorMessage("Sorry! We didn't find any content :(", 404)
	}
	return
}

func MapHandler(c f.Context) {
	var ok bool
	var lang, input string
	if lang, ok = c.FormValue("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
		return
	}

	if input, ok = c.FormValue("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
		return
	}

	if uri := getMap(input, lang); uri != "" {
		c.JsonWrite(map[string]string{"geojson": uri}, 200)
	} else {
		c.WriteErrorMessage("Sorry! We didn't find any content :(", 404)
	}
	return
}
