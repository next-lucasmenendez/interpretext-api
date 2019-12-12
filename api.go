package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tokenizer "github.com/next-lucasmenendez/interpretext-tokenizer"
	langdetector "github.com/next-lucasmenendez/interpretext-lang-detector"
    keywords "github.com/next-lucasmenendez/interpretext-keyword-extractor"
	postagger "github.com/next-lucasmenendez/interpretext-postagger"
	summarizer "github.com/next-lucasmenendez/interpretext-text-summarizer"
)

func checkInput(w http.ResponseWriter, r *http.Request) (input string) {
	if input = r.FormValue("input"); input == "" {
		http.Error(w, "No input text provided", 400)
		log.Panic("No input")
	}

	return
}

func responseJson(w http.ResponseWriter, d interface{}) {
	if res, e := json.Marshal(d); e != nil {
		http.Error(w, "Error parsing JSON response.", 500)
		log.Panic(e)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	var input string = checkInput(w, r)

	d := map[string]interface{} {"lang": langdetector.Suggest(input)}
	responseJson(w, d)
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	var (
		tokens [][]string
		input string = checkInput(w, r)
		sentences []string = tokenizer.Sentences(input)
	)

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

func keywordsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		results [][]string
		tokens []string
		input string = checkInput(w, r)
		lang string = langdetector.Suggest(input)
		sentences []string = tokenizer.Sentences(input)
	)

	if len(sentences) > 0 {
		for _, s := range sentences {
			tokens = append(tokens, tokenizer.Words(s)...)
		}
	} else {
		tokens = append(tokens, tokenizer.Words(input)...)
	}

	var e error
	if results, e = keywords.GetTags(tokens, lang, 12); e != nil {
        http.Error(w, "Error extracting keywords from input text.", 500)
		log.Fatal(e.Error())
    }

	d := map[string]interface{} {"keywords": results}
	responseJson(w, d)
}

func postaggingHandler(w http.ResponseWriter, r *http.Request) {
	var models string = os.Getenv("MODELS")
	if models == "" {
		http.Error(w, "No models configured for PoS tagging.", 500)
		log.Fatal("Empty MODELS env variable")
	}

	var (
		input string = checkInput(w, r)
		lang string = langdetector.Suggest(input)
	)

	var tokens []string
	if sentences := tokenizer.Sentences(input); len(sentences) > 0 {
		for _, s := range sentences {
			tokens = append(tokens, tokenizer.Words(s)...)
		}
	} else {
		tokens = append(tokens, tokenizer.Words(input)...)
	}

	var modelPath string = fmt.Sprintf("%s/%s", models, lang)
	if model, e := postagger.LoadModel(modelPath); e != nil {
		http.Error(w, "Error loading model for PoS tagging.", 500)
		log.Fatal(e)
	} else {
		var (
			tagger *postagger.Tagger = postagger.NewTagger(model)
			tagged [][]string = tagger.Tag(tokens)
		)

		if len(tagged) == 0 {
			http.Error(w, "No tags were found.", 404)
		}

		d := map[string]interface{} { "tagged": tagged }
		responseJson(w, d)
	}
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	var (
		input string = checkInput(w, r)
		lang string = langdetector.Suggest(input)
	)

	if s, e := summarizer.NewText(input, lang); e != nil {
		http.Error(w, "Error analyzing input text.", 500)
		log.Fatal(e.Error())
	} else {
		d := map[string]interface{} { "summary": s.Summarize() }
		responseJson(w, d)
	}
}

func startApi() {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	http.HandleFunc("/keywords", keywordsHandler)
	http.HandleFunc("/postagging", postaggingHandler)
	http.HandleFunc("/summary", summaryHandler)

	fmt.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
