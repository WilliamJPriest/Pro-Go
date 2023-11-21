package database

import (
	"fmt"
	"log"
)

func GetArticles() (string, error) {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal("%w", err)
	}

	defer db.Close()
	_, err = db.Query("Select * from Articles")
	if err != nil{
		return "", fmt.Errorf("couldn't find in articles table: %w", err)
	}
	return "", nil
}