package database

import (
	"fmt"
	"log"

	"github.com/williamjPriest/HTMXGO/models"
)

func GetArticles() (error) {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal("%w", err)
	}
	var authorNames []models.ArticleData
	defer db.Close()
	rows, err := db.Query("Select Author from Articles")
	if err != nil{
		return fmt.Errorf("couldn't find in articles table: %w", err)
	}
	defer rows.Close()
	var authorName models.ArticleData
	for rows.Next(){
		err := rows.Scan(&authorName.Author)
		if err != nil{
			return fmt.Errorf("failed to scan row: %w", err)
		}
		authorNames = append(authorNames, authorName)
		
	}
	fmt.Println(authorNames)

	return nil
}