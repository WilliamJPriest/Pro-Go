package database

import (
	"fmt"
)



func CheckBookMarks(title string) (error) {

	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	fmt.Println(title)
	//also check that logged user also matches db user
	row := db.QueryRow("SELECT title FROM Bookmarks WHERE title = $1", title)
    
    if err := row.Scan(); err != nil {
        return fmt.Errorf("%w", err)
    }

	return nil
}

func AddBookMarks(author string, title string, desc string,urltoimage string,content string) {
	fmt.Println(author + " added")

}

func RemovedBookMarks(author string, title string, desc string,urltoimage string,content string) {
	fmt.Println(author +" removed")

}