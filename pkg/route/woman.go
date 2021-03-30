package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func Selection(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	selectedWemen := make([]int, 0, 5)
	// HTTPメソッド確認
	if r.Method != "POST" {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	//
	s := r.FormValue("selected_wemen")
	formData := strings.Split(s, ",")
	log.Println(formData)
	log.Println(len(formData))
	log.Printf("%T", formData)
	if len(formData) != 5 {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	for _, v := range formData {
		num, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		selectedWemen = append(selectedWemen, num)
	}
	res, err := json.Marshal(selectedWemen)
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
