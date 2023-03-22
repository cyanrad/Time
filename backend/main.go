package main

import (
	"log"
	"os"

	"main/api"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatalln(err)
	}

	s := api.NewServer(db)
	err = s.Start(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
