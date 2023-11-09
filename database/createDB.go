package database

import (
	"log"


	_ "github.com/lib/pq"
)

func Create(){
	db,err := ConnectToDB()
	if err != nil{
		log.Fatal("%w", err)
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
