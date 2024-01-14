package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}


	DBlink := os.Getenv("DB_LINK")
	dsn := DBlink
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w", err)
		
	}
	return db, nil
}