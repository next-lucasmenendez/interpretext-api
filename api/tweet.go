package api

import (
	a "github.com/lucasmenendez/gobstract"
	"strings"
	"regexp"
	"fmt"
)

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
	
	return fmt.Sprintf("\"%s\"", tweet)
}