package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var dbCon *sql.DB

func init() {
	// DBとBackendを接続
	dbCon = NewDB()
	// actressテーブルにデータ挿入
	actressesFile := "./pkg/db/actress_id.txt"
	names, err := getActressesDataFromText(actressesFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = insertActresses(names)
	if err != nil {
		log.Fatalln(err)
	}
}

// Create a connection with MySQL
func NewDB() *sql.DB {
	dBUser := os.Getenv("DB_USER")
	dBPass := os.Getenv("DB_PASS")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", dBUser, dBPass, dbProtocol, dbEndpoint, dbName)
	dbCon, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return dbCon
}

// Actressテーブルの初期データ挿入
func insertActresses(names []string) error {
	// トランザクション開始
	tx, err := dbCon.Begin()
	if err != nil {
		return err
	}

	// クエリ定義
	query := `INSERT INTO actresses (name) VALUES (?)`

	// 141人分トランザクションでInsertする
	for _, name := range names {
		tx.Exec(query, name)
	}

	// コミット
	err = tx.Commit()
	// エラー発生したらロールバックする
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	} else {
		return nil
	}
}

// actress_id.txtから名前を配列にしたものを取得
func getActressesDataFromText(fileName string) (names []string, err error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return names, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		actress := strings.Split(line, ":")
		name := actress[1]
		names = append(names, name)
	}
	return names, nil
}
