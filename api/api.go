package api

import (
	"strconv"
	f "github.com/lucasmenendez/framework.go"
)

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
	var err error

	var ok bool
	var lang, input, raw_count string
	if lang, ok = c.FormValue("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
	}

	if input, ok = c.FormValue("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
	}

	var count int
	if raw_count, ok = c.FormValue("count"); !ok {
		count = 10
	} else if count, err = strconv.Atoi(raw_count); err != nil {
		c.WriteErrorMessage("Bad count provided.", 400)
		return
	}

	var data map[string]interface{} = getSummary(input, lang, count)
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
