package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/souhub/avzeus-backend/pkg/model"
)

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
