package db

import (
	"log"
	"main/models"

	"github.com/jmoiron/sqlx"
)

// func clearPeriodsTable(db *sqlx.DB) {
// 	const clear = `TRUNCATE TABLE periods;`

// 	db.MustExec(clear)
//}

func ReadPeriods(db *sqlx.DB, dateStamp int32) (periods []models.Period) {
	err := db.Select(&periods, "SELECT * FROM periods WHERE date_stamp = $1 ORDER BY location ASC;", dateStamp)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// func createPeriods(db *sqlx.DB, p []models.Period) (periods []models.Period) {
// 	// Generating query string
// 	queryString := "INSERT INTO periods (date_stamp, location, duration, type, title, description) VALUES "
// 	for _, period := range p {
// 		queryString += fmt.Sprintf("(%v, %v, %v, '%v', '%v', '%v'),", period.DateStamp, period.Location, period.Duration, period.Type, period.Title, period.Description)
// 	}
// 	queryString = queryString[:len(queryString)-1]
// 	queryString += " RETURNING *;"
// 	fmt.Println(queryString)

// 	// generating the query
// 	rows, err := db.Queryx(queryString)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// this is so fucking dumb but what can i do
// 	for rows.Next() {
// 		var period models.Period
// 		err := rows.StructScan(&period)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		periods = append(periods, period)
// 	}
// 	return
// }

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
