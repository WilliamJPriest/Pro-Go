package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func Create(){
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

	_, err = db.Exec("CREATE TABLE if not exists Users (UserID SERIAL PRIMARY KEY, username char(100),password char(100));")
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	_, err = db.Exec("CREATE TABLE if not exists Bookmarks (BookmarkID SERIAL PRIMARY KEY, username char(100),author TEXT, title TEXT, Description TEXT, Url TEXT, UrltoImage Text, Content TEXT);")
	if err != nil {
		log.Fatal("failed to execute query", err)
	}


}
