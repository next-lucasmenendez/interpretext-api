package main


import (
	a "github.com/lucasmenendez/gobstract"
	f "github.com/lucasmenendez/framework.go"
	"strconv"
	"strings"
	"regexp"
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
		var sentence string = abstract.GetBestSentence()
		var keywords []string = abstract.GetKeywords()

		var splitter *regexp.Regexp = regexp.MustCompile(`\s`)
		var words []string = splitter.Split(sentence, -1)

		var bestSentence string
		for _, word := range words {
			for _, keyword := range keywords {
				if strings.TrimSpace(strings.ToLower(word)) == strings.TrimSpace(strings.ToLower(keyword)) {
					word = fmt.Sprintf("#%s", word)
				}
			}

			bestSentence = fmt.Sprintf("%s %s", bestSentence, word)
		}

		c.JsonWrite(map[string]string{"best_sentence": bestSentence}, 200)
	}
	return
}

func main() {
	s := f.New()

	s.DebugMode(true)

	port_raw := os.Getenv("PORT")
	if port, err := strconv.Atoi(port_raw); err != nil {
		fmt.Println("No port provided. Using default :9999")
		port = 9999
	} else {
		s.SetPort(port)
	}

	s.POST("/", handler, auth)
	s.Run()
}