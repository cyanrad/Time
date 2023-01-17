package main

import (
	"log"

	"main/api"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const serverAddress = "127.0.0.1:8080"

func main() {
	// Running the database
	db, err := sqlx.Connect("postgres", "user=root password=secret dbname=main sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	s := api.NewServer(db)
	err = s.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
