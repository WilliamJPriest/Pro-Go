package database

import (
	"fmt"
)



func CheckBookMarks(title string, username string) (error) {

	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}

	//also check that logged user also matches db user
	row := db.QueryRow("SELECT title, username FROM Bookmarks WHERE title = $1 AND username = $2", title, username)

    if err := row.Scan(&title, &username); err != nil {
        return fmt.Errorf("%w", err)
    }

	return nil
}

func AddBookMarks(author string, title string, desc string,urltoimage string,content string, username string) error{
	fmt.Println(author + " added to this user:"+ username)
	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}

	_, err = db.Exec("INSERT INTO Bookmarks (author, title, Description, UrlToImage, Content, username) VALUES ($1, $2,$3,$4,$5,$6)", author, title, desc, urltoimage, content, username)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}


func RemovedBookMarks(title string,  username string) error{
	fmt.Println(title +" removed from this user"+ username)
	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}

	_, err = db.Exec("DELETE FROM Bookmarks WHERE title = $1 AND username = $2", title, username)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil

}