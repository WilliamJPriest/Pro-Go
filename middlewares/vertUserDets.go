package middlewares

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"
	"github.com/williamjPriest/HTMXGO/utils"
	"github.com/joho/godotenv"
)

func VerifyUser(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("username")
		email := req.PostFormValue("email")
		validEmail := utils.ValidateEmail(email)
		if !validEmail{
			t := template.Must(template.ParseGlob("templates/register-error.html"))
			t.Execute(w, nil)
			return

		}
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
		defer db.Close()
		rows, err := db.Query("SELECT * FROM Users WHERE username = $1 or email = $2 ", username, email)
		if err != nil {
			log.Fatal("Error executing SQL query: %w", err)
		}
		defer rows.Close()
	
		if rows.Next() {
			t := template.Must(template.ParseGlob("templates/register-error.html"))
			t.Execute(w, nil)
		} else {
			endpointHandler(w, req)
		}

	})
}