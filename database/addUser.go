package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func AddUser(username string, password string){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	DBlink := os.Getenv("DB_LINK")
	dsn := DBlink
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	_, err = db.Exec("INSERT INTO Users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
}
