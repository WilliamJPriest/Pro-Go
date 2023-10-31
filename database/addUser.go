package database

import (
	"database/sql"
	"fmt"

	// "log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)



func AddUser(username string, password []byte) (bool,error){
	err := godotenv.Load()
	if err != nil {
		return false,  fmt.Errorf("Error loading .env file: %w", err)
	 
	}
	DBlink := os.Getenv("DB_LINK")
	dsn := DBlink
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return false,  fmt.Errorf("Failed to execute query: %w" ,err)
	}

	_, err = db.Exec("INSERT INTO Users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return true, nil
}
