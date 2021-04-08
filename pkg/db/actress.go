package db

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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

// 初期データを actresses テーブルに代入
func initializeActresses() (err error) {
	// sqlファイルからクエリ作成＋actresses テーブル作成
	query := parseSqlFile("actresses")
	_, err = dbCon.Exec(query)
	if err != nil {
		err = errors.New("Failed to create actresses table")
		return err
	}

	// actressesテーブルのレコードが空ならトランザクションでデータ挿入
	if isEmpty("actresses") {
		// トランザクション開始
		tx, err := dbCon.Begin()
		if err != nil {
			err = errors.New("Failed to bigin a transaction")
			return err
		}

		// actresses_id.txt から名前を抽出
		names := getActressesDataFromText("./actress_id.txt")

		// 抽出した名前を全て actresses テーブルに insert
		query = `INSERT INTO actresses (name) VALUES (?)`
		for _, name := range names {
			_, err = tx.Exec(query, name)
			if err != nil {
				err = errors.New("Failed to Insert actresses")
				tx.Rollback()
			}
		}

		// コミット
		err = tx.Commit()
		if err != nil {
			err = errors.New("Failed to commit the transaction in actresses table")
			err = tx.Rollback()
			return err
		}
	}
	// actressesテーブルが空じゃないなら終了
	return nil
}

// actress_id.txtから名前を配列にしたものを取得
func getActressesDataFromText(fileName string) (names []string) {
	filePath := fmt.Sprintf("./seeds/%s", fileName)
	fp, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to get actresses data from %s", fileName)
		err = errors.New(msg)
		log.Fatalln(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		actress := strings.Split(line, ":")
		name := actress[1]
		names = append(names, name)
	}
	return names
}
