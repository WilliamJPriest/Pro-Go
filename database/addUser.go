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

	_, err = db.Exec("INSERT INTO Users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return fmt.Errorf("failed to execute query3: %w", err)
	}
	return nil
}
