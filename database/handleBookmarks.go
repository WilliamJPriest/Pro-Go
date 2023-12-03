package database

import (
	"fmt"
	"github.com/williamjPriest/HTMXGO/models"
)



func CheckBookMarks(title string, username string) (error) {

	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	row := db.QueryRow("SELECT title, username FROM Bookmarks WHERE title = $1 AND username = $2", title, username)
	defer db.Close()
	
    if err := row.Scan(&title, &username); err != nil {
        return fmt.Errorf("%w", err)
    }
	return nil
}

func AddBookMarks(author string, title string, desc string,url string,urltoimage string, username string) error{
	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	defer db.Close()
	fmt.Println(url)

	_, err = db.Exec("INSERT INTO Bookmarks (author, title, description, url, urlToImage, username) VALUES ($1, $2,$3,$4,$5,$6)", author, title, desc,url, urltoimage, username)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}


func RemovedBookMarks(title string,  username string) error{
	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Bookmarks WHERE title = $1 AND username = $2", title, username)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil

}

func GetBookMarks(username string) ([]models.BookmarkData, error){
	db,err := ConnectToDB()
	defer db.Close()
	var bookmarks []models.BookmarkData
	rows, err := db.Query("SELECT author, title, description, url, urlToImage FROM Bookmarks where username= $1 order by BookmarkID desc", username)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookmark models.BookmarkData
		err := rows.Scan(&bookmark.Author,&bookmark.Title,&bookmark.Description,&bookmark.Url,&bookmark.UrlToImage)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		bookmarks = append(bookmarks, bookmark)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}


	
	return bookmarks, nil
}