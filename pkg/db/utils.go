package db

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/souhub/avzeus-backend/pkg/model"
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
	filePath := fmt.Sprintf("./pkg/db/sql/%s.sql", fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("%sのSQLファイルの読み込みに失敗", fileName)
		err = errors.New(msg)
		log.Fatalln(err)
	}
	query := fmt.Sprintf("%s", content)
	return query
}

// wemen_id.txtから名前と画像パスをWoman構造体に代入したものを取得
func getWemenDataFromText(fileName string) (wemen model.Wemen) {
	filePath := fmt.Sprintf("./pkg/db/seeds/%s", fileName)
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
		params := strings.Split(line, " ")
		woman := model.Woman{
			Name:      params[0],
			ImagePath: params[1],
		}
		wemen = append(wemen, woman)
	}
	return wemen
}
