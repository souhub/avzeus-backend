package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

// Actressを全県取得し、構造体に入れて返す
func FetchActresses() (actresses model.Actresses) {
	rows := fetchActressesRows()
	for rows.Next() {
		var actress model.Actress
		err := rows.Scan(&actress.ID, &actress.Name, &actress.ImagePath)
		if err != nil {
			log.Println(err)
		}
		actresses = append(actresses, actress)
	}
	return actresses
}

// AIから受け取ったrecommended_actresses_idsをもとに、レコードを取得し、構造体に入れて返す
func FetchRecommendedActresses(ids []int) (recommendedActresses model.Actresses, err error) {
	for _, id := range ids {
		row, err := fetchActressRow(id)
		if err != nil {
			log.Println(err)
		}
		var actress model.Actress
		err = row.Scan(&actress.ID, &actress.Name, &actress.ImagePath)
		if err != nil {
			log.Println(err)
		}
		recommendedActresses = append(recommendedActresses, actress)
	}
	return recommendedActresses, err
}

// Actressテーブルからレコードを全取得
func fetchActressesRows() *sql.Rows {
	query := `SELECT * FROM actresses`
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

// 任意のレコードを１つActressテーブルから取得
func fetchActressRow(id int) (row *sql.Row, err error) {
	query := `SELECT * FROM actresses WHERE id=?`
	row = dbCon.QueryRow(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("No rows")
		}
	}
	return row, err
}

// // Actressテーブルの初期データ挿入
// func insertActresses(names []string) error {
// 	// トランザクション開始
// 	tx, err := dbCon.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	// クエリ定義
// 	query := `INSERT INTO actresses (name) VALUES (?)`

// 	// 141人分トランザクションでInsertする
// 	for _, name := range names {
// 		tx.Exec(query, name)
// 	}

// 	// コミット
// 	err = tx.Commit()
// 	// エラー発生したらロールバックする
// 	if err != nil {
// 		if err = tx.Rollback(); err != nil {
// 			return err
// 		}
// 		return err
// 	} else {
// 		return nil
// 	}
// }

// actress_id.txtから名前を配列にしたものを取得
// func getActressesDataFromText(fileName string) (names []string, err error) {
// 	fp, err := os.Open(fileName)
// 	if err != nil {
// 		return names, err
// 	}
// 	defer fp.Close()

// 	scanner := bufio.NewScanner(fp)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		actress := strings.Split(line, ":")
// 		name := actress[1]
// 		names = append(names, name)
// 	}
// 	return names, nil
// }

// actressテーブルにレコードが存在するか否か
// func isExist() bool {
// 	query := `SELECT COUNT(*) FROM actresses`
// 	rows, err := dbCon.Query(query)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var count int
// 	for rows.Next() {
// 		if err := rows.Scan(&count); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	if count == 0 {
// 		return false
// 	}
// 	return true
// }
