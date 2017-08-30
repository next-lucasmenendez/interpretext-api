package api

import (
	f "github.com/lucasmenendez/framework.go"
	"strconv"
)

func TweetHandler(c f.Context) {
	var err error
	var form f.Form
	if form, err = c.ParseForm(); err != nil {
		c.WriteError(err, 500)
	}

	var ok bool
	var lang, input string
	if lang, ok = form.Get("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
		return
	}

	if input, ok = form.Get("input"); !ok {
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
	var form f.Form
	if form, err = c.ParseForm(); err != nil {
		c.WriteError(err, 500)
	}

	var ok bool
	var lang, input, raw_count string
	if lang, ok = form.Get("lang"); !ok {
		c.WriteErrorMessage("No language provided.", 400)
		return
	}

	if input, ok = form.Get("input"); !ok {
		c.WriteErrorMessage("No text provided.", 400)
		return
	}

	var count int
	if raw_count, ok = form.Get("count"); !ok {
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
