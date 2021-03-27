package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/souhub/av-zeus/pkg/db"
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
	w.Write(res)
}
