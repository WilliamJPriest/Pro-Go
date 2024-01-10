package database

import (
	"database/sql"
	"fmt"
	"os"


)

func ConnectToDB() (*sql.DB, error){

	DBlink := os.Getenv("DB_LINK")
	dsn := DBlink
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w", err)
		
	}
	return db, nil
}