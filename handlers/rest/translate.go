package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"sources/m1/translation"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		println("default language")
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	log.Print(word)
	log.Print("language: " + language)
	translation := translation.Translation(word, language)
	if translation == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")

	}
}
