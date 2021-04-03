package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gorilla/schema"
	"github.com/souhub/avzeus-backend/pkg/db"
	"github.com/souhub/avzeus-backend/pkg/model"
)

const BaseURL = "http://localhost:8080"

// GET
// /actress
func Actresses(w http.ResponseWriter, r *http.Request) {
	actresses := db.FetchActresses()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Charset", "utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(w).Encode(actresses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// GET
// /wemen
func Wemen(w http.ResponseWriter, r *http.Request) {
	wemen := db.FetchWemen()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Charset", "utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(w).Encode(wemen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// POST
// /selection
func PostDataToAI(w http.ResponseWriter, r *http.Request) {
	// HTTPメソッド確認
	if r.Method != "POST" {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	// formのデータを解析し、受け取る
	r.ParseForm()
	// string型で受け取る
	selectedWemenDataStr := r.FormValue("selected_wemen")
	// string型を[]stringに変換し、選択された人数を確認する準備
	formData := strings.Split(selectedWemenDataStr, ",")
	// 5人選択されてなければリダイレクトさせる
	if len(formData) != 5 {
		http.Redirect(w, r, "http://localhost:8080/selection", 301)
		return
	}
	// リダイレクト URIを作成
	url := fmt.Sprintf("http://localhost:5000/inputted-data/%s", selectedWemenDataStr)
	// リダイレクト
	http.Redirect(w, r, url, 301)
}

var decoder = schema.NewDecoder()

// GET
// /outputted-data
func GetDataFromAI(w http.ResponseWriter, r *http.Request) {
	// クエリを抽出
	requestedQuery := r.URL.Query()
	// TrainingData 構造体にクエリを入れる
	var trainingData model.TrainingData
	err := decoder.Decode(&trainingData, requestedQuery)
	if err != nil {
		log.Fatalln(err)
	}
	// Training data の states と epsilons を保存してidを返す
	trainingDataID, err := db.InsertTraining(trainingData)
	if err != nil {
		log.Fatalln(err)
	}
	// パースして変換 *url.URL に型変換
	redirectURL, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	// リクエストを投げるパスを指定
	requestPath := "recommendation"
	// パスの結合がある場合はここで結合される
	redirectURL.Path = path.Join(redirectURL.Path, requestPath)
	// クエリをセット
	query := redirectURL.Query()
	query.Set("ids", trainingData.ActressesIDs)
	query.Add("id", fmt.Sprint(trainingDataID))

	query.Add("states", trainingData.States)     // 見る用に一時的に追加
	query.Add("epsilons", trainingData.Epsilons) // 見る用に一時的に追加

	redirectURL.RawQuery = query.Encode()
	// リダイレクト
	http.Redirect(w, r, redirectURL.String(), 302)
}

// GET
// /recommendation
func GetRecommendedActresses(w http.ResponseWriter, r *http.Request) {
	// クエリを抽出
	requestedQuery := r.URL.Query()
	// q := requestedQuery.Get("id")
	qq := requestedQuery.Get("ids")
	// []int 型に変換
	recommendedActressesIDs, err := convertStrToIntArray(qq)
	if err != nil {
		log.Fatalln(err)
	}
	// クエリの配列をもとに、DBからデータ取得
	recommendedActresses, err := db.FetchRecommendedActresses(recommendedActressesIDs)
	if err != nil {
		log.Fatalln(err)
	}
	// jsonとして書き込む
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Charset", "utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(recommendedActresses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
