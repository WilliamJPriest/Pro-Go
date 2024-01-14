package database

import (
	"fmt"

	_ "github.com/lib/pq"
)



func AddUser(username string, password []byte) (error){
	db,err := ConnectToDB()
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (username, email, password) VALUES ($1, $2,$3)", username, password)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
