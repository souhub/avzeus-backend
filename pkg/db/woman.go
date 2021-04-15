package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

// 初期データを wemen テーブルに代入
func initializeWemen() (err error) {
	// sqlファイルからクエリ作成＋wemenテーブル作成
	query := parseSqlFile("woman/create_table")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create wemen table")
		return err
	}
	// wemenテーブルのレコードが空なら挿入
	if isEmpty("wemen") {
		// トランザクション開始
		tx, err := dbCon.Begin()
		if err != nil {
			err = errors.New("Failed to bigin a transaction")
			return err
		}

		// woman_id.txt から wemen デーブルのデータを抽出
		wemen := getWemenDataFromText("woman_id.txt")

		// 抽出したデータを wemen テーブルに挿入
		query := `INSERT INTO wemen (name, image_path) VALUES (?, ?)`
		for _, woman := range wemen {
			_, err = tx.Exec(query, woman.Name, woman.ImagePath)
			if err != nil {
				err = errors.New("Failed to Insert wemen")
				tx.Rollback()
			}
		}

		// コミット
		err = tx.Commit()
		if err != nil {
			err = errors.New("Failed to commit the transaction in wemen table")
			err = tx.Rollback()
			return err
		}
	}
	return nil
}

// Fetch all wemen from database
func FetchWemen() (wemen model.Wemen) {
	rows := fetchWemenRows()
	for rows.Next() {
		var woman model.Woman
		err := rows.Scan(&woman.ID, &woman.Name, &woman.ImagePath)
		if err != nil {
			log.Fatalln(err)
		}
		wemen = append(wemen, woman)
	}
	return wemen
}

// Fetch all rows from wemen table
func fetchWemenRows() *sql.Rows {
	query := parseSqlFile("woman/select_wemen")
	rows, err := dbCon.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows")
		} else {
			log.Println(err)
		}
	}
	return rows
}

func resetWemen() (err error) {
	query := `DELETE FROM wemen`
	_, err = dbCon.Exec(query)
	return err
}
