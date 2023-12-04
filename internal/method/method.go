package method

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexandrMasalov/kanji/internal/connectdb"
)

type responseKanji = map[string]string

func GetKanji(w http.ResponseWriter, r *http.Request) {

	kanji, err := connectdb.GetDBValue()
	if err != nil {
		fmt.Println(err)
	}
	text := responseKanji{
		"stringKanji": kanji,
	}
	json.NewEncoder(w).Encode(text)
}
