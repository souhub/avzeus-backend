package route

import (
	"bytes"
	"encoding/json"
	"errors"
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
// /recommendation
func GetRecommendationActresses(w http.ResponseWriter, r *http.Request) {
	// HTTPメソッド確認
	log.Println(r.Method)
	if r.Method != "POST" {
		endpoint := FrontendURL + "/selection"
		http.Redirect(w, r, endpoint, http.StatusMethodNotAllowed)
		return
	}
	// クエリを取得
	// requestedQuery := r.URL.Query()
	// selectedWemenDataStr := requestedQuery.Get("selected_wemen_ids")

	// リクエストボディをパース
	reqbody, err := ioutil.ReadAll(r.Body)
	// field名は小文字だとjsonタグを付けてもアクセスできないためパース失敗する
	// struct名大文字でも小文字でもどちらでもいい
	type requestedData struct {
		SelectedWemenDataStr string `json:"selected_wemen_ids"`
	}
	var reqData requestedData
	err = json.Unmarshal(reqbody, &reqData)
	if err != nil {
		err = errors.New("Failed to read a body of a recommendation reques")
		log.Println(err)
		return
	}
	selectedWemenDataStr := reqData.SelectedWemenDataStr

	//  フォームで5人選択されているかチェック
	checkSelectionForm(selectedWemenDataStr, w, r)

	// AIサーバーに投げるリクエストURLを作成
	target := "recommendation"
	targeURL, err := url.Parse(AIURL)
	if err != nil {
		err = errors.New("Failed to create a recommendation request for AI")
		log.Fatalln(err)
	}
	targeURL.Path = path.Join(targeURL.Path, target)
	q := targeURL.Query()
	q.Set("selected_wemen_ids", selectedWemenDataStr)
	targeURL.RawQuery = q.Encode()

	// AIサーバーにリクエストを投げて JSON を TrainingData 構造体で受ける
	resp, err := http.Get(targeURL.String())
	if err != nil {
		err = errors.New("Failed to get a response from AI server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Failed to read a response from AI server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var trainingData model.TrainingData
	err = json.Unmarshal(body, &trainingData)
	if err != nil {
		err = errors.New("Failed to parse the recommended actresses data from AI server")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Trainingテーブルにstatesとepsilonsを保存＋ID返却
	id, err := db.InsertTrainingData(trainingData)
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

	// idとrecommended_actressesを1つの構造体にまとめる
	type RecommendedData struct {
		RecommendedActresses model.Actresses `json:"recommended_actresses"`
		ID                   int             `json:"id"`
	}

	data := RecommendedData{
		RecommendedActresses: recommended_actresses,
		ID:                   id,
	}

	// フロントエンドサーバーにJSONで返す
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// POST
// /result
func Result(w http.ResponseWriter, r *http.Request) {
	/// Validate a http method
	if r.Method != "POST" {
		http.Error(w, "This method isn't allowed for a GetResult handler", http.StatusMethodNotAllowed)
		return
	}
	// リクエストボディをResult構造体に入れてパース
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.New("Failed to read a result")
		log.Println(err)
		return
	}
	var result model.Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		err = errors.New("Failed to parse a result to the Result struct")
		log.Println(err)
		return
	}
	// resultsテーブルにtraining_idが登録されているか
	// 登録されていれば終了
	if db.IsResultExists(result.TrainingID) {
		return
	}
	// 登録されていなければresultをDBに保存
	if err = db.InsertResult(result); err != nil {
		err = errors.New("Failed to insert a result to the Result struct")
		log.Println(err)
		return
	}
}

func Training(trainingIDs []int) {
	var trainingDatas []model.TrainingData
	// training dataを配列で取得
	for _, id := range trainingIDs {
		// training_idが一致するstatesを取得
		statesArr, err := db.FetchVectors("states", id)
		if err != nil {
			err = errors.New("Failed to fetch states from the db")
			log.Println(err)
			return
		}
		// training_idが一致するepsilonsを取得
		epsilonsArr, err := db.FetchVectors("epsilons", id)
		if err != nil {
			err = errors.New("Failed to fetch epsilons from the db")
			log.Println(err)
			return
		}
		// Training構造体に代入してJSONでAIにPOST
		trainingData := model.TrainingData{
			ID:       id,
			States:   statesArr,
			Epsilons: epsilonsArr,
		}
		trainingDatas = append(trainingDatas, trainingData)
	}
	// training dataを解析
	jsonTrainingDatas, err := json.Marshal(trainingDatas)
	if err != nil {
		err = errors.New("Failed to convert TrainingDatas struct to JSON")
		log.Println(err)
		return
	}
	// POSTのボディ用意
	postBody := bytes.NewBuffer(jsonTrainingDatas)
	endpoint := AIURL + "/training"
	// AIサーバーにPOSTリクエスト送信
	http.Post(endpoint, "application/json", postBody)
	log.Println(jsonTrainingDatas)
}
