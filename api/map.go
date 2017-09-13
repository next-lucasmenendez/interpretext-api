package api

import (
	"fmt"
	"strings"
	"net/http"
	gcd "github.com/beeva-lucasmenendez/gocountrydetect"
	"io/ioutil"
	"encoding/json"
)

const API string = "https://country-geojson-api.herokuapp.com"

func getMap(text, lang string) (map[string]interface{}, error) {
	var countries gcd.Countries = gcd.Detect(text, lang)
	var names []string
	for _, c := range countries {
		names = append(names, strings.ToLower(c.Name))
	}

	if len(countries) == 0 {
		return nil, nil
	}
	var plainNames string = strings.Join(names, ",")

	var err error
	var res *http.Response
	var uri string = fmt.Sprintf("%s/search/%s?lang=%s", API, plainNames, lang)
	if res, err = http.Get(uri); err != nil {
		return nil, err
	}

	var bodyBytes []byte
	defer res.Body.Close()

    if bodyBytes, err = ioutil.ReadAll(res.Body); err != nil {
        return nil, err
    }

    var data map[string]interface{}
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
        return nil, err
    }

    return data, nil
}

