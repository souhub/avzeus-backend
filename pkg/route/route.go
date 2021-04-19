package route

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/google/uuid"
	"github.com/souhub/avzeus-backend/pkg/db"
	"github.com/souhub/avzeus-backend/pkg/dmm"
	"github.com/souhub/avzeus-backend/pkg/model"
)

// GET
// /actress
func Actresses(w http.ResponseWriter, r *http.Request) {
	actresses := db.FetchActresses()
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
	err := json.NewEncoder(w).Encode(wemen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// POST
// /recommendation
func Recommendation(w http.ResponseWriter, r *http.Request) {
	// HTTPメソッド確認
	if r.Method != "POST" {
		endpoint := FrontendURL + "/selection"
		http.Redirect(w, r, endpoint, http.StatusTemporaryRedirect)
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
	trainingID, err := db.InsertTrainingData(trainingData)
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
	// 推薦された女優データにDMMのデータを追加
	recommended_actresses, err = dmm.AddDataToActresses(recommended_actresses)
	if err != nil {
		log.Println(err)
	}

	// idとrecommended_actressesを1つの構造体にまとめる
	type RecommendedData struct {
		RecommendedActresses model.Actresses `json:"recommended_actresses"`
		TrainingID           int             `json:"training_id"`
	}

	data := RecommendedData{
		RecommendedActresses: recommended_actresses,
		TrainingID:           trainingID,
	}

	// フロントエンドサーバーにアクセスさせるエンドポイントにJSONでおいていく
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// POST
// /image-clipping
func ImageClipping(w http.ResponseWriter, r *http.Request) {
	// メソッド確認
	if r.Method != "POST" {
		err := errors.New("Uplading an image allow only POST method")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	// フォームから画像を受け取る
	// FormFileは必要に応じて自動的にParseMultipleFormかParseFromを呼び出す
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		msg := fmt.Sprintf("Failed to receive an image from image-uploader form. %s", err)
		err = errors.New(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// ファイル名が一意になるようにUUIDを用いて変更
	// MTCNN(画像切り抜きAI)は拡張子がない保存パスを受け取るとエラーになるためここで対応しておく
	fileName := fileHeader.Filename
	ext := path.Ext(fileName)
	if ext == "" {
		log.Fatalln("No extension")
	}
	fileName = uuid.NewString() + ext

	// imagePath := fmt.Sprintf("tmp/%s", fileName)
	imagePath := fileName
	// S3にアップロード
	if err = s3Upload(imagePath, file); err != nil {
		log.Fatalln(err)
	}
	log.Println("Upload a file to S3 was completed.")
	// AIサーバーにPOSTするJSONの準備
	type ImageData struct {
		ImagePath string `json:"image_path"`
	}
	// 一時保存場所の絶対パスを作る
	imageURL, err := createURL(S3URL, nil, fileName)
	if err != nil {
		log.Fatalln(err)
	}
	imageData := ImageData{
		// S3のオブジェクトキー"がtmp/[filename]"と表されるからimagePathをAIサーバーに渡す
		ImagePath: imagePath,
	}
	jsonImageData, err := json.Marshal(imageData)
	if err != nil {
		msg := fmt.Sprintf("Failed to marshal an image path to JSON. %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	// POSTのボディ用意
	postBody := bytes.NewBuffer(jsonImageData)
	endpoint, err := createURL(AIURL, nil, "image-clipping")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Preparing for a request to AI was completed.")

	// AIサーバーにPOSTリクエスト送信
	resp, err := http.Post(endpoint, "application/json", postBody)
	if err != nil {
		log.Println(err)
		log.Println("The AI was failed to learn because of a request error")
	}
	defer resp.Body.Close()

	log.Println("Sending a request to AI was completed.")
	// レスポンスボディを解析
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type RespParser struct {
		IDs []int `json:"ids"`
	}
	respParser := new(RespParser)
	if err = json.Unmarshal(respBody, respParser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ids := respParser.IDs
	idsStr := ""
	for _, id := range ids {
		idStr := fmt.Sprintf("%d ", id)
		idsStr += idStr
	}

	log.Println("Parsing a response from AI was completed.")

	// 画像の一時保存場所のURLと、AIから受け取ったID配列とをURLに乗せてフロントエンドにリダイレクト
	queries := map[string]string{
		"ids":        idsStr,
		"image_path": imageURL,
	}
	if queries["ids"] == "" {
		err = s3Delete(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		redirectURL, err := createURL(FrontendURL, nil, "image-uploader")
		if err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	}
	redirectURL, err := createURL(FrontendURL, queries, "image-uploader")
	if err != nil {
		log.Fatalln(err)
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
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
