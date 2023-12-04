package server

import (
	"log"
	"net/http"

	"github.com/AlexandrMasalov/kanji/internal/method"
)

func Run() {
	http.HandleFunc("/get-kanji", method.GetKanji)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
