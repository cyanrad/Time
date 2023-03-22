package db

import (
	"log"
	"main/models"

	"github.com/jmoiron/sqlx"
)

func ReadPeriods(db *sqlx.DB, dateStamp int32) (periods []models.Period) {
	err := db.Select(&periods, "SELECT * FROM periods WHERE date_stamp = $1 ORDER BY location ASC;", dateStamp)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreatePeriod(db *sqlx.DB, p models.Period) (period models.Period) {
	row, err := db.NamedQuery("INSERT INTO periods (date_stamp, location, duration, type, title, description) VALUES (:date_stamp, :location, :duration, :type, :title, :description) RETURNING *;", p)
	if err != nil {
		log.Fatalln(err)
	}

	if row.Next() {
		err = row.StructScan(&period)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return
}

func UpdatePeriod(db *sqlx.DB, p models.Period) (period models.Period) {
	row, err := db.NamedQuery("UPDATE periods SET title = :title, description = :description, type = :type, duration = :duration WHERE id=:id  RETURNING *;", p)
	if err != nil {
		log.Fatalln(err)
	}

	if row.Next() {
		err = row.StructScan(&period)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return
}

func DeletePeriod(db *sqlx.DB, ID int64) {
	db.MustExec("DELETE FROM periods WHERE id=$1;", ID)
}
