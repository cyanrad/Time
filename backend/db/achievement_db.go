package db

import (
	"log"
	"main/models"

	"github.com/jmoiron/sqlx"
)

func ReadAchievements(db *sqlx.DB, dateStamp int32) (achievements []models.Achievement) {
	err := db.Select(&achievements, "SELECT * FROM achievements WHERE date_stamp = $1 ORDER BY id ASC;", dateStamp)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateAchievement(db *sqlx.DB, a models.Achievement) (achievement models.Achievement) {
	row, err := db.NamedQuery("INSERT INTO achievements (date_stamp, text) VALUES (:date_stamp, :text) RETURNING *;", a)
	if err != nil {
		log.Fatalln(err)
	}

	if row.Next() {
		err = row.StructScan(&achievement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return
}

func UpdateAchievement(db *sqlx.DB, a models.Achievement) (achievement models.Achievement) {
	row, err := db.NamedQuery("UPDATE achievements SET text = :text WHERE id=:id  RETURNING *;", a)
	if err != nil {
		log.Fatalln(err)
	}

	if row.Next() {
		err = row.StructScan(&achievement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return
}

func DeleteAchievement(db *sqlx.DB, ID int64) {
	db.MustExec("DELETE FROM achievements WHERE id=$1;", ID)
}
