package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/souhub/avzeus-backend/pkg/db"
)

// GET
// /wemen
func Wemen(w http.ResponseWriter, r *http.Request) {
	wemen := db.FetchWemen()
	res, err := json.Marshal(wemen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Charset", "utf-8")
	w.Write(res)
}
