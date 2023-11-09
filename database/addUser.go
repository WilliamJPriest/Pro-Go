package database

import (
	"database/sql"
	"fmt"


	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)



func AddUser(username string, password []byte) (error){
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file: %w", err)
	 
	}
	DBlink := os.Getenv("DB_LINK")
	dsn := DBlink
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("Failed to execute query2: %w" ,err)
	}

	_, err = db.Exec("INSERT INTO Users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return fmt.Errorf("failed to execute query3: %w", err)
	}
	return nil
}
