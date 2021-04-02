package db

import (
	"errors"

	"github.com/souhub/avzeus-backend/pkg/model"
)

func InsertTraining(trainingData model.TrainingData) (id int, err error) {
	// トランザクション開始
	tx, err := dbCon.Begin()
	if err != nil {
		return id, err
	}
	// INSERT クエリ実行
	insertQuery := `INSERT INTO training (states, epsilons)
				  	VALUES (?, ?)`
	tx.Exec(insertQuery, trainingData.States, trainingData.Epsilons)
	// SELECT クエリ実行
	selectQuery := `SELECT LAST_INSERT_ID()`
	tx.QueryRow(selectQuery)
	// コミット
	err = tx.Commit()
	// エラー発生したらロールバックする
	if err != nil {
		if err = tx.Rollback(); err != nil {
			err = errors.New("Failed to rollback")
		}
		return id, err
	} else {
		return id, nil
	}
}
