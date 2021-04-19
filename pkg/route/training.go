package route

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/souhub/avzeus-backend/pkg/db"
	"github.com/souhub/avzeus-backend/pkg/model"
)

// 1週間に1回学習させる(AI学習用)
func init() {
	go learn()
}

// ゴルーチンさせる関数(AI学習用)
func learn() {
	for {
		// 1週間待つ
		time.Sleep(604800 * time.Second)
		// 実行
		postTrainingData()
	}
}

// TrainingデータをAIにPOST(AI学習用)
func postTrainingData() {
	// 1週間分のTrainingデータIDを配列で取得
	trainingIDs, err := db.FetchTrainingIDsForOneWeek()
	if err != nil {
		log.Fatalln(err)
	}
	// ID配列をもとにTraining IDと一致するstate,epsilon取得
	var trainingDatas []model.TrainingData
	for _, id := range trainingIDs {
		// training_idが一致するvalをresultsテーブルから取得
		// errが起きた場合、resultsのval存在しない＝どの女優リンクもクリックされてないのでAIにデータを渡さない
		resultVal, err := db.FetchResult(id)
		if err != nil {
			log.Println(err)
			err = errors.New("Failed to fetch result from the db")
			continue
		}
		log.Println(resultVal)
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
			// ID:       id,
			States:   statesArr,
			Epsilons: epsilonsArr,
			Result:   resultVal,
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
	endpoint := AIURL + "/learning"
	// trainingDatasがnilの状態でAIサーバーにリクエストを送るとAIサーバーが落ちるから
	if len(trainingDatas) == 0 {
		msg := `{"msg":Failed: The data didn't send to AI. There are now row in results table "status: -2}`
		log.Printf("%s", msg)
		return
	}
	// AIサーバーにPOSTリクエスト送信
	resp, err := http.Post(endpoint, "application/json", postBody)
	if err != nil {
		log.Println(err)
		log.Println("The AI was failed to learn because of a request error")
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Failed to read a response of learning result from AI server")
		log.Println(err)
		return
	}
	log.Printf("%s", respBody)
}
