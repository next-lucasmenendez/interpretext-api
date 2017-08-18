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

func getTweet(input, lang string) string {
	if abstract, err := a.NewAbstract(input, lang); err != nil {
		return ""
	} else {
		var sentence string = abstract.GetBestSentence()
		var keywords []string = abstract.GetKeywords()

		var bestSentence string = composeTweet(sentence, keywords)
		return cleanTweet(bestSentence)
	}
	return ""
}

func composeTweet(tweet string, keywords []string) string {
	var splitter *regexp.Regexp = regexp.MustCompile(`\s`)
	var rawWords []string = splitter.Split(tweet, -1)

	var words []string = make([]string, len(rawWords))
	for _, word := range rawWords {
		for _, keyword := range keywords {
			if strings.TrimSpace(strings.ToLower(word)) == strings.TrimSpace(strings.ToLower(keyword)) {
				word = fmt.Sprintf("#%s", word)
				break
			}
		}
		words = append(words, word)
	}

	return strings.Join(rawWords, " ")
}

func cleanTweet(tweet string) string {
	var cleaner *regexp.Regexp = regexp.MustCompile(`(\[.+]|\(.+\))`)
	var clean string = cleaner.ReplaceAllString(tweet, "")
	tweet = strings.TrimSpace(clean)

	if len(tweet) > 113 {
		tweet = fmt.Sprintf("\"%s...\"", tweet[:110])
	} else {
		tweet = fmt.Sprintf("\"%s\"", tweet)
	}
	return tweet
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

	var tweet string = getTweet(input, lang)
	if tweet != "" {
		c.JsonWrite(map[string]string{"best_sentence": tweet}, 200)
	} else {
		c.WriteErrorMessage("Tweet not found :(", 404)
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