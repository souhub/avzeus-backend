package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

// Trainingテーブルの初期化
func initializeStates() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training/create_table_states")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create states table")
		return err
	}
	return nil
}

// Statesテーブルの初期化
func initializeEpsilons() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training/create_table_epsilons")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create epsilons table")
		return err
	}
	return nil
}

// Epsilonsテーブルの初期化
func initializeResults() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training/create_table_results")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create results table")
		return err
	}
	return nil
}

// Resultテーブルの初期化
func initializeTraining() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training/create_table_training")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create training table")
		return err
	}
	return nil
}

// Training dataの挿入
func InsertTrainingData(trainingData model.TrainingData) (id int, err error) {
	// trainingテーブルにレコードを作り、そのIDを取得
	id, err = getTrainingID()
	if err != nil {
		return id, err
	}
	// AIサーバーから受け取ったベクトル配列をStates,Epsilons構造体に入れる
	statesArr := trainingData.States
	epsilonsArr := trainingData.Epsilons
	states := convertArrToStruct(statesArr, id)
	if err != nil {
		err = errors.New("Failed to convert states arr to states struct")
		return id, err
	}
	epsilons := convertArrToStruct(epsilonsArr, id)
	if err != nil {
		err = errors.New("Failed to convert epailons arr to epailons struct")
		return id, err
	}
	// データベースに登録
	if err = insertVectors(states, "state"); err != nil {
		return id, err
	}
	if err = insertVectors(epsilons, "epsilon"); err != nil {
		err = errors.New("Failed to insert epsilons to database")
		return id, err
	}
	return id, nil
}

// AIサーバーから受け取ったベクトル配列をStates,Epsilons構造体に入れる
func convertArrToStruct(arr []float64, trainingID int) (vector model.Vectors) {
	var vectors model.Vectors
	for _, val := range arr {
		vector := model.Vector{
			Val:        val,
			TrainingID: trainingID,
		}
		vectors = append(vectors, vector)
	}
	return vectors
}

// Training IDを取得
func getTrainingID() (id int, err error) {
	// トランザクション開始
	tx, err := dbCon.Begin()
	if err != nil {
		return id, err
	}
	// 空のINSERT クエリ実行
	insertQuery := `INSERT INTO training () VALUES ()`
	_, err = tx.Exec(insertQuery)
	if err != nil {
		tx.Rollback()
		err = errors.New("Training data の挿入失敗")
		return id, err
	}
	// SELECT クエリ実行し、IDを返す
	selectQuery := parseSqlFile("training/select_last_insert_id")
	row := tx.QueryRow(selectQuery)
	// レコードをスキャンして返ってきたIDを取得
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		err = errors.New("Training data のID 取得失敗")
		return id, err
	}
	// コミット
	err = tx.Commit()
	// エラー発生したらロールバックする
	if err != nil {
		err = errors.New("Failed to commit a training data")
		if err = tx.Rollback(); err != nil {
			err = errors.New("Failed to rollback")
		}
		return id, err
	}
	return id, nil
}

// Vector(States,Epsilons)のデータベースへの挿入
func insertVectors(vectors model.Vectors, typeName string) (err error) {
	// トランザクション開始
	tx, err := dbCon.Begin()
	if err != nil {
		msg := fmt.Sprintf("Failed to begin a transaction in %s", typeName)
		err = errors.New(msg)
		return err
	}
	// クエリ定義
	filePath := fmt.Sprintf("training/insert_%s", typeName)
	query := parseSqlFile(filePath)
	// statesテーブルにデータ挿入
	for _, vector := range vectors {
		_, err = tx.Exec(query, vector.Val, vector.TrainingID)
		if err != nil {
			// log.Fatal(err)
			tx.Rollback()
			msg := fmt.Sprintf("Failed to insert a %s", typeName)
			err = errors.New(msg)
			return err
		}
	}
	// コミット
	err = tx.Commit()
	if err != nil {
		msg := fmt.Sprintf("Failed to commit the transaction in %s", typeName)
		err = errors.New(msg)
		return err
	}
	return nil
}

// 選択されたactressIDを挿入
func InsertResult(result model.Result) (err error) {
	query := parseSqlFile("training/insert_selected_actress_id")
	_, err = dbCon.Exec(query, result.Val, result.TrainingID)
	if err != nil {
		err = errors.New("Failed to insert a selected actress id")
		return err
	}
	return nil
}

// training_idが一致するstatesまたはepsilonsを全て取得
func FetchVectors(name string, trainingID int) (vectorsArr []float64, err error) {
	filePath := fmt.Sprintf("training/select_%s", name)
	query := parseSqlFile(filePath)
	rows, err := dbCon.Query(query, trainingID)
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("Failed to get %s from the database", name)
		err = errors.New(msg)
		return vectorsArr, err
	}
	for rows.Next() {
		var vector model.Vector
		err = rows.Scan(&vector.Val)
		if err != nil {
			log.Println(err)
		}
		vectorsArr = append(vectorsArr, vector.Val)
	}
	return vectorsArr, nil
}

// training_idが一致するresultを取得
func FetchResult(trainingID int) (val int, err error) {
	filePath := "training/select_result"
	query := parseSqlFile(filePath)
	row := dbCon.QueryRow(query, trainingID)
	if err = row.Scan(&val); err != nil {
		err = errors.New("Failed to get val from results table")
		return val, err
	}
	return val, err
}

// 1週間分のrainingテーブルのIDを返す（AI学習用）
func FetchTrainingIDsForOneWeek() (ids []int, err error) {
	query := parseSqlFile("training/select_training_for_one_week")
	rows, err := dbCon.Query(query)
	if err != nil {
		return ids, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// resultsテーブルに同じtraining_idがあるか否か
func IsResultExists(trainingID int) bool {
	query := `SELECT id FROM results WHERE training_id=? LIMIT 1`
	row := dbCon.QueryRow(query, trainingID)
	var id int
	if err := row.Scan(&id); err != nil {
		log.Printf("%s. This is a validation of result table,NOT an error", err)
		log.Println(err)
		return false
	}
	return true
}
