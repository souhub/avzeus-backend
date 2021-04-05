package db

import (
	"database/sql"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

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
	query := `SELECT * FROM wemen`
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
