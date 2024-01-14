package routes

import (

	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginPageHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/login.html"))
	t.Execute(w, nil)
}	



func LoginHandler(w http.ResponseWriter, req *http.Request){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	secretCode := os.Getenv("SECRET_CODE")
	var SecretKey = []byte(secretCode)
	Username := req.PostFormValue("username")


	token := jwt.New(jwt.SigningMethodHS256)
	expiration := time.Now().Add(6000 * time.Minute)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expiration.Unix()
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatalf("failed to login %s", err)
	}
	

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration,
		HttpOnly: true,
		Secure:   true,
	})

	t := template.Must(template.ParseGlob("templates/welcome.html"))
	t.Execute(w, Username)



}

func GuestLoginHandler(w http.ResponseWriter, req *http.Request){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	secretCode := os.Getenv("SECRET_CODE")
	var SecretKey = []byte(secretCode)

	Username := "Guest101"


	token := jwt.New(jwt.SigningMethodHS256)
	expiration := time.Now().Add(60 * time.Minute)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expiration.Unix()
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatalf("failed to login %s", err)
	}
	

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration,
		HttpOnly: true,
		Secure:   true,
	})

	t := template.Must(template.ParseGlob("templates/welcome.html"))
	t.Execute(w, Username)



}

func LogoutHandler(w http.ResponseWriter, req *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Expires: time.Now(),

	})
	w.Header().Set("Hx-Redirect", "/")
	

}

func RegisterPageHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/register.html"))
	t.Execute(w, nil)
	

}

func RegisterHandler(w http.ResponseWriter, req *http.Request){
	username := req.PostFormValue("username")
	email := req.PostFormValue("email")
	password := req.PostFormValue("password")
	
	bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
	if err != nil{
		log.Fatalf("failed to hash: %s", err)
	}

	err = database.AddUser(username,email,bcrypt)
	if err != nil{
		log.Fatalf("failed to add user: %s", err)
	}
	
	t := template.Must(template.ParseGlob("templates/login.html"))
	t.Execute(w, nil)
}