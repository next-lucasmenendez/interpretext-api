package main

import (
	"encoding/json"
	"fmt"
	"github.com/beeva-labs/lang-detector"
	"github.com/beeva-labs/text-summarizer"
	"github.com/beeva-labs/text-tokenizer"
	"log"
	"net/http"
	"os"
)

func responseJson(w http.ResponseWriter, d interface{}) {
	if res, e := json.Marshal(d); e != nil {
		log.Panic(e)
	} else {
		w.Write(res)
	}
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	var e error
	var input string
	if input = r.FormValue("input"); input == "" {
		log.Panic("No input")
	}

	var lang string = langdetector.Suggest(input)
	var s *summarizer.Text
	if s, e = summarizer.NewText(input, lang); e != nil {
		log.Panic(e)
	}

	d := map[string]interface{} { "summary": s.Summarize() }
	responseJson(w, d)
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	var input string
	if input = r.FormValue("input"); input == "" {
		log.Panic("No input")
	}

	d := map[string]interface{} {"lang": langdetector.Suggest(input)}
	responseJson(w, d)
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	var input string
	if input = r.FormValue("input"); input == "" {
		log.Panic("No input")
	}

	var tokens [][]string
	var sentences []string = tokenizer.Sentences(input)
	if len(sentences) > 0 {
		for _, s := range sentences {
			tokens = append(tokens, tokenizer.Words(s))
		}
	} else {
		tokens = append(tokens, tokenizer.Words(input))
	}

	d := map[string]interface{} {"tokens": tokens}
	responseJson(w, d)
}

func startApi() {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	http.HandleFunc("/summary", summaryHandler)
	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
