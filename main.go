package main


import (
	a "github.com/lucasmenendez/gobstract"
	f "github.com/lucasmenendez/framework.go"
	"strconv"
	"fmt"
	"os"
)

func auth(c f.Context) {
	fmt.Println("Middleware")
	c.Continue()
}

func handler(c f.Context) {
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

	if abstract, err := a.NewAbstract(input, lang); err != nil {
		c.WriteError(err, 500)
	} else {
		var res map[string]interface{} = map[string]interface{} {
			"highlights": abstract.GetHightlights(),
			"keywords": abstract.GetKeywords(),
			"best_sentence": abstract.GetBestSentence(),
		}

		c.JsonWrite(res, 200)
	}
	return
}

func main() {
	s := f.New()

	s.DebugMode(true)

	port_raw := os.Getenv("PORT")
	if port, err := strconv.Atoi(port_raw); err != nil {
		fmt.Println("No port provided.")
		return
	} else {
		s.SetPort(port)
	}

	s.POST("/", handler, auth)
	s.Run()
}