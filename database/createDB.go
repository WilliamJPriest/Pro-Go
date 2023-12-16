package database

import (
	"log"


	_ "github.com/lib/pq"
)

func Create() error{
	db,err := ConnectToDB()
	if err != nil{
		log.Fatal("%w", err)
	}

	_, err = db.Exec("CREATE TABLE if not exists Users (UserID SERIAL PRIMARY KEY, username char(100),password char(100),email char(100));")
	if err != nil {
		log.Fatal("failed to execute Users query", err)
	}

	// _, err = db.Exec("ALTER TABLE Users ADD COLUMN ;")
	// if err != nil {
	// 	log.Fatal("failed to execute alter query", err)
	// }

	_, err = db.Exec("CREATE TABLE if not exists Bookmarks (BookmarkID SERIAL PRIMARY KEY, username char(100),author tEXT, title TEXT, description TEXT, Url TEXT, urlToImage Text, content TEXT);")
	if err != nil {
		log.Fatal("failed to execute Bookmarks query", err)
	}

	_, err = db.Exec("CREATE TABLE if not exists Articles (ArticlesID SERIAL PRIMARY KEY, author TEXT, title TEXT, description TEXT, Url TEXT, urlToImage Text, content TEXT,likes INTEGER, dislikes INTEGER, rank TEXT, Badge Text);")
	if err != nil {
		log.Fatal("failed to execute Articles query", err)
	}


	return nil

}
