package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLとBackendを接続
var dbCon = NewDB()

func init() {
	// DBのセットアップ完了まで初期化処理を待たせる
	for {
		err := dbCon.Ping()
		if err != nil {
			log.Println("Waiting for setting up a database...")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	// 初期化実行
	initializeDB()
	log.Println("DB initialization was completed")
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
	err = initializeStates()
	if err != nil {
		log.Fatalln(err)
	}
	err = initializeEpsilons()
	if err != nil {
		log.Fatalln(err)
	}
	err = initializeResults()
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
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dBUser, dBPass, dbProtocol, dbEndpoint, dbPort, dbName)
	dbCon, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return dbCon
}
