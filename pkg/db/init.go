package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLとBackendを接続
var dbCon = NewDB()

func init() {
	// 初期化実行
	initializeDB()
}

// データベースの初期化
func initializeDB() {
	err := initializeActresses()
	if err != nil {
		log.Fatalln(err)
	}
	err = initializeWemen()
	if err != nil {
		log.Fatalln(err)
	}
	err = initializeTraining()
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
