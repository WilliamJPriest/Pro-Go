package middlewares

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func VertifyLogin(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")

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
		var storedPasswordHash string
		err = db.QueryRow("SELECT password FROM Users WHERE username = $1", username).Scan(&storedPasswordHash)
		if err != nil {
			log.Fatal("Error executing SQL query: %w", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
		if err != nil {
			t := template.Must(template.ParseFiles("register-error.html"))
			t.Execute(w, nil)
		}

	})
}