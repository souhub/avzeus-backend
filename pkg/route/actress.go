package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/souhub/avzeus-backend/pkg/db"
)

// GET
// /actress
func Actresses(w http.ResponseWriter, r *http.Request) {
	actresses := db.FetchActresses()
	res, err := json.Marshal(actresses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Charset", "utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func Recommendation(w http.ResponseWriter, r *http.Request) {
	url := "https://connpass.com/api/v1/event/?keyword=python&count=1"
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer res.Body.Close()
	http.Redirect(w, r, "http://localhost:880/recommendation", 301)
}
