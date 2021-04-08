package db

import (
	"errors"

	"github.com/souhub/avzeus-backend/pkg/model"
)

// Trainingテーブルの初期化
func initializeTraining() (err error) {
	// sqlファイルからクエリ作成＋trainingテーブル作成
	query := parseSqlFile("training")
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
	insertQuery := `INSERT INTO training (states, epsilons)
				  	VALUES (?, ?)`
	_, err = tx.Exec(insertQuery, trainingData.States, trainingData.Epsilons)
	if err != nil {
		tx.Rollback()
		err = errors.New("Training data の挿入失敗")
		return id, err
	}
	// SELECT クエリ実行
	selectQuery := `SELECT LAST_INSERT_ID()`
	row := tx.QueryRow(selectQuery)
	// レコードをスキャンして返ってきたIDを代入
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
