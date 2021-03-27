package db

import (
	"database/sql"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

func FetchActresses() (actresses model.Actresses) {
	rows := fetchActressesRows()
	for rows.Next() {
		var actress model.Actress
		err := rows.Scan(&actress.ID, &actress.Name, &actress.ImagePath, &actress.Vector)
		if err != nil {
			log.Fatalln(err)
		}
		actresses = append(actresses, actress)
	}
	return actresses
}

func fetchActressesRows() *sql.Rows {
	query := `SELECT * FROM actresses`
	rows, err := dbCon.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Panicln("No rows")
		} else {
			log.Println(err)
		}
	}
	return rows
}
