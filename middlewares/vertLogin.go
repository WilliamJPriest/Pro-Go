package middlewares

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func VertifyLogin(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")

		err, encypted = CompareHashAndPassword(SQLPassword, password)

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file: %w", err)
		 
		}
		DBlink := os.Getenv("DB_LINK")
		dsn := DBlink
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Failed to execute query: %w" ,err)
		}
		//use username to fetch password
		// decrypt and see if it matches user input
		//if it does go to main page
		//if not send error html, password didn't match
		rows, err := db.Query("SELECT * FROM Users WHERE username = $1, password = $2", username, password)
		if err != nil {
			log.Fatal("Error executing SQL query: %w", err)
		}
		defer rows.Close()
		err, encypted = CompareHashAndPassword(rows.password, password)

	})
}