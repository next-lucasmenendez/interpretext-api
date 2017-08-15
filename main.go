package main


import (
	a "github.com/lucasmenendez/gobstract"
	f "github.com/lucasmenendez/framework.go"
)

func auth(c f.Context) {
	c.Continue()
}

func hightlights(c f.Context) {
	var err error
	var form f.Form
	if form, err = c.ParseForm(); err != nil {
		c.WriteError(err, 500)
	}

	lang := form["lang"]
	input := form["text"]

	if abstract, err := a.NewAbstract(input, lang); err != nil {
		c.WriteError(err, 500)
	} else {
		var res map[string][]string = map[string][]string {
			"highlights": abstract.GetHightlights(),
		}

		c.JsonWrite(res, 200)
	}
}

func keywords(c f.Context) {
	var err error
	var form f.Form
	if form, err = c.ParseForm(); err != nil {
		c.WriteError(err, 500)
	}

	lang := form["lang"]
	input := form["text"]

	if abstract, err := a.NewAbstract(input, lang); err != nil {
		c.WriteError(err, 500)
	} else {
		var res map[string][]string = map[string][]string {
			"keywords": abstract.GetKeywords(),
		}

		c.JsonWrite(res, 200)
	}
}

func bestSentence(c f.Context) {
	var err error
	var form f.Form
	if form, err = c.ParseForm(); err != nil {
		c.WriteError(err, 500)
	}

	lang := form["lang"]
	input := form["text"]

	if abstract, err := a.NewAbstract(input, lang); err != nil {
		c.WriteError(err, 500)
	} else {
		var res map[string]string = map[string]string {
			"best_sentence": abstract.GetBestSentence(),
		}

		c.JsonWrite(res, 200)
	}
}

func main() {
	s := f.New()

	s.DebugMode(true)
	s.SetPort(9999)

	s.POST("/highlights", hightlights)
	s.POST("/keywords", keywords)
	s.POST("/best-sentence", bestSentence)
	s.Run()
}