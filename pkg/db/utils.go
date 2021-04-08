package db

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// 任意のテーブルが空かどうか
func isEmpty(tableName string) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	row := dbCon.QueryRow(query)
	var cnt int
	err := row.Scan(&cnt)
	if err != nil {
		log.Fatalln(err)
	}
	if cnt == 0 {
		return true
	}
	return false
}

// SQLファイルを読み込んでクエリ文字列を返す
func parseSqlFile(fileName string) string {
	filePath := fmt.Sprintf("./sql/%s.sql", fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("%sのSQLファイルの読み込みに失敗", fileName)
		err = errors.New(msg)
		log.Fatalln(err)
	}
	query := fmt.Sprintf("%s", content)
	return query
}
