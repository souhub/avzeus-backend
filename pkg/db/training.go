package db

import (
	"errors"

	"github.com/souhub/avzeus-backend/pkg/model"
)

// Trainingテーブルの初期化
func initializeTraining() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training/create_table")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create training table")
		return err
	}
	return nil
}

// Training data の挿入
func InsertTraining(trainingData model.TrainingData) (id int, err error) {
	// トランザクション開始
	tx, err := dbCon.Begin()
	if err != nil {
		return id, err
	}
	// INSERT クエリ実行
	insertQuery := parseSqlFile("training/insert_training")
	_, err = tx.Exec(insertQuery, trainingData.States, trainingData.Epsilons)
	if err != nil {
		tx.Rollback()
		err = errors.New("Training data の挿入失敗")
		return id, err
	}
	// SELECT クエリ実行
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

// 選択されたactressIDを挿入
func InsertSelectedActressID(id, selectedID int) (err error) {
	query := parseSqlFile("training/insert_selected_actress_id")
	_, err = dbCon.Exec(query, selectedID, id)
	if err != nil {
		err = errors.New("Failed to insert selected actress id")
		return err
	}
	return nil
}
