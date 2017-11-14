package gotokenizer

import "regexp"

func Sentences(s string) (sentences []string) {
	var (
		numP       = regexp.MustCompile(`([0-9]+)\.([0-9]+)`)
		quoutesP   = regexp.MustCompile(`("|'|“|”|’|«|»)`)
		pstopsP    = regexp.MustCompile(`"(.+)\.(.+)"`)
		revpstopsP = regexp.MustCompile(`{partial_stop}`)
		stopsP     = regexp.MustCompile(`[^..][!?.]\s`)
		resP       = regexp.MustCompile(`\*\|\*`)
		dotP       = regexp.MustCompile(`{stop}`)

		noNum     = numP.ReplaceAllString(s, `$1*|*$2`)
		noQuoutes = quoutesP.ReplaceAllString(noNum, "\"")
		noPstops  = pstopsP.ReplaceAllString(noQuoutes, "\"$1{partial_stop}$2\"")
		noStops   = stopsP.ReplaceAllString(noPstops, `$0{stop}`)
		text      = resP.ReplaceAllString(noStops, `.`)
		resText   = revpstopsP.ReplaceAllString(text, `.`)
	)

	sentences = dotP.Split(resText, -1)
	return sentences
}

func Words(s string) (tokens []string) {
	var (
		ws = regexp.MustCompile(`\s|\t`)
		cr = regexp.MustCompile(`\s|\t|"|'|\.|,|\:`)

		temp = ws.Split(s, -1)
	)

	for _, w := range temp {
		var _w string = cr.ReplaceAllLiteralString(w, "")
		if _w != "" {
			tokens = append(tokens, w)
		}
	}

	return tokens
}
