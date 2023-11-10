package database

import (
	"fmt"
)



func CheckBookMarks(title string, username string) (error) {

	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	fmt.Println(title)
	fmt.Println(username)
	//also check that logged user also matches db user
	row := db.QueryRow("SELECT title, username FROM Bookmarks WHERE title = $1, username = $2", title, username)
    
    if err := row.Scan(); err != nil {
        return fmt.Errorf("%w", err)
    }

	return nil
}

func AddBookMarks(author string, title string, desc string,urltoimage string,content string, username string) {
	fmt.Println(author + " added to this user:"+ username)

}

func RemovedBookMarks(author string, title string, desc string,urltoimage string,content string, username string) {
	fmt.Println(author +" removed from this user"+ username)

}