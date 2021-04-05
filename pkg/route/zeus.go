package route

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/souhub/avzeus-backend/pkg/db"
	"github.com/souhub/avzeus-backend/pkg/model"
)

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
func GetRecommendationActresses(w http.ResponseWriter, r *http.Request) {
	// HTTPメソッド確認
	if r.Method != "GET" {
		endpoint := BackendURL + "/selection"
		http.Redirect(w, r, endpoint, 301)
		return
	}
	// クエリを取得
	requestedQuery := r.URL.Query()
	selectedWemenDataStr := requestedQuery.Get("selected_wemen_ids")

	//  フォームで5人選択されているかチェック
	checkSelectionForm(selectedWemenDataStr, w, r)

	// AIサーバーに投げるリクエストURLを作成
	target := "recommendation"
	targeURL, err := url.Parse(AIURL)
	if err != nil {
		log.Println("zeus.go line:73")
		log.Fatalln(err)
	}
	targeURL.Path = path.Join(targeURL.Path, target)
	q := targeURL.Query()
	q.Set("selected_wemen_ids", selectedWemenDataStr)
	targeURL.RawQuery = q.Encode()

	// AIサーバーにリクエストを投げて JSON を TrainingData 構造体で受ける
	resp, err := http.Get(targeURL.String())
	if err != nil {
		log.Println("zeus.go line:84")
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("zeus.go line:92")
		log.Fatalln(err)
	}
	var trainingData model.TrainingData
	err = json.Unmarshal(body, &trainingData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// DB から女優データ持ってくる
	actresses_ids := trainingData.RecommendedActressesIDs
	recommended_actresses, err := db.FetchRecommendedActresses(actresses_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// フロントエンドサーバーにJSONで返す
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(recommended_actresses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// GET
// /outputted-data
// func GetDataFromAI(w http.ResponseWriter, r *http.Request) {
// 	// クエリを抽出
// 	requestedQuery := r.URL.Query()
// 	// TrainingData 構造体にクエリを入れる
// 	var trainingData model.TrainingData
// 	err := decoder.Decode(&trainingData, requestedQuery)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// Training data の states と epsilons を保存してidを返す
// 	trainingDataID, err := db.InsertTraining(trainingData)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// パースして変換 *url.URL に型変換
// 	redirectURL, err := url.Parse(BaseURL)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// リクエストを投げるパスを指定
// 	requestPath := "recommendation"
// 	// パスの結合がある場合はここで結合される
// 	redirectURL.Path = path.Join(redirectURL.Path, requestPath)
// 	// クエリをセット
// 	query := redirectURL.Query()
// 	query.Set("actresses_ids", trainingData.ActressesIDs)
// 	query.Add("training_id", fmt.Sprint(trainingDataID))

// 	query.Add("states", trainingData.States)     // 見る用に一時的に追加
// 	query.Add("epsilons", trainingData.Epsilons) // 見る用に一時的に追加

// 	redirectURL.RawQuery = query.Encode()
// 	// リダイレクト
// 	http.Redirect(w, r, redirectURL.String(), 302)
// }

// GET
// /recommendation
// func GetRecommendedActresses(w http.ResponseWriter, r *http.Request) {
// 	// クエリを抽出
// 	requestedQuery := r.URL.Query()
// 	// q := requestedQuery.Get("id")
// 	qq := requestedQuery.Get("ids")
// 	// []int 型に変換
// 	recommendedActressesIDs, err := convertStrToIntArray(qq)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// クエリの配列をもとに、DBからデータ取得
// 	recommendedActresses, err := db.FetchRecommendedActresses(recommendedActressesIDs)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// jsonとして書き込む
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Accept-Charset", "utf-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	err = json.NewEncoder(w).Encode(recommendedActresses)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Println(err)
// 		return
// 	}
// }
