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

// POST
// /selection
func Selection(w http.ResponseWriter, r *http.Request) {
	// HTTPメソッド確認
	if r.Method != "POST" {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	// formのデータを解析し、受け取る
	r.ParseForm()
	selectedWemen := make([]int, 0, 5)
	// string型で受け取る
	s := r.FormValue("selected_wemen")
	// string型を[]stringに変換
	formData := strings.Split(s, ",")
	// 5人選択されてなければリダイレクトさせる
	if len(formData) != 5 {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	// []string型から[]int型に変換
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
	http.Redirect(w, r, "/results", 200)
	// http.Redirect(w, r, "http://localhost:880/recommendation", 301)
}

func passDataToAI(w http.ResponseWriter, res []byte) {

}

// GET
//
