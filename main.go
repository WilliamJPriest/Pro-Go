package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)




type userData struct{
	Username string
	Password string
}

type ArticlesData struct{
	Articles []ArticleData `json:"articles"`
}


type ArticleData struct{
	Author string `json:"author"`
	Title string  `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	UrlToImage string `json:"urlToImage"`
}

var MUserName string
var MPassword []byte

var SecretKey = []byte("SecretYouShouldHide")


func main(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	ApiKey := os.Getenv("API_KEY")
  
	MainPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		res, err := http.Get(ApiKey)
		if err != nil{
			log.Fatalf("response issue: %s", err)
		}
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject ArticlesData

		json.Unmarshal(responseData, &responseObject)

		fmt.Println(responseObject.Articles[0].Author)
		t.Execute(w, nil)
	}	
	loginPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("login.html"))
		t.Execute(w, nil)
	}	

	registerPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("register.html"))
		t.Execute(w, nil)

	}

	loginHandler := func(w http.ResponseWriter, req *http.Request){
		Username := req.PostFormValue("username")
		// Password := req.PostFormValue("password")
		// if Username != MUserName {
		// 	log.Fatalf("This name didn't match: %s", Username)
  
		// }
		// err := bcrypt.CompareHashAndPassword([]byte(MPassword) , []byte(Password))
	    // if err != nil{
		// 	log.Fatalf("didn't match: %s", err)
		// }

		generateJWT(Username)

		
		

	}
	registerHandler := func(w http.ResponseWriter, req *http.Request){
		Username := req.PostFormValue("username")
		Password := req.PostFormValue("password")
		bcrypt,err := bcrypt.GenerateFromPassword([]byte(Password), 5  )
		if err != nil{
			log.Fatalf("failed to hash: %s", err)
		}
		MUserName= Username
		MPassword=bcrypt
		//once registered there should switch to the login page.

	}


	http.HandleFunc("/",MainPageHandler)
	http.HandleFunc("/loginPage",loginPageHandler)
	http.HandleFunc("/registerPage", registerPageHandler )	
	http.HandleFunc("/login", loginHandler )
	http.HandleFunc("/register", registerHandler )

	log.Fatal(http.ListenAndServe(":8000",nil))
}

func generateJWT(Username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	
	fmt.Println(tokenString)


 return tokenString, nil

}