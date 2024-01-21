package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(w http.ResponseWriter, req *http.Request){
	Username := req.PostFormValue("username")
	setJWtToken(w, Username)

}

func GuestLoginController(w http.ResponseWriter, req *http.Request){
	Username := "Guest101"
	setJWtToken(w, Username)

}

func setJWtToken(w http.ResponseWriter, Username string){
	err := godotenv.Load()
	if err != nil {
	log.Printf("Error loading .env file")
	}
	secretCode := os.Getenv("SECRET_CODE")
	var SecretKey = []byte(secretCode)


	token := jwt.New(jwt.SigningMethodHS256)
	expiration := time.Now().Add(6000 * time.Minute)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expiration.Unix()
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("failed to login %s", err)
	}
	

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration,
		HttpOnly: true,
		Secure:   true,
	})

}

func RegisterController(w http.ResponseWriter, req *http.Request){
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	
	bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
	if err != nil{
		log.Printf("failed to hash: %s", err)
	}

	err = database.AddUser(username,bcrypt)
	if err != nil{
		log.Printf("failed to add user: %s", err)
	}

}