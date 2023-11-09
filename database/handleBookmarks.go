package database

import (
	"fmt"
)

func AddBookMarks() {

}

func CheckBookMarks(title string) bool {

	db,err := ConnectToDB()
	if err != nil{
		fmt.Println("%w",err)
		return false
	}
	fmt.Println(title)
	row := db.QueryRow("SELECT title FROM Bookmarks WHERE title = $1", title)
    
    if err := row.Scan(); err != nil {
        fmt.Println("%w", err)
        return false
    }

	return true
}