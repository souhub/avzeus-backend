package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbCon *sql.DB

func init() {
	dbCon = NewDB()
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
