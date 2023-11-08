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
	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/models"
	"github.com/williamjPriest/HTMXGO/middlewares"
)




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
			fmt.Println(err)
		}
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		var responseObject models.ArticlesData

		json.Unmarshal(responseData, &responseObject)

		// just testing something
		_, err = req.Cookie("token")
		if err != nil {
			responseObject.IsLoggedIn = false
			t.Execute(w, responseObject)
		}else{
			responseObject.IsLoggedIn = true
			t.Execute(w, responseObject)
		}
		
		
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
		username := req.PostFormValue("username")


		token := jwt.New(jwt.SigningMethodHS256)
		expiration := time.Now().Add(10 * time.Minute)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = expiration.Unix()
		claims["authorized"] = true
		claims["user"] = username
	
		tokenString, err := token.SignedString(models.SecretKey)
		if err != nil {
			log.Fatalf("failed to login %s", err)
		}
		
		fmt.Println(tokenString)
	
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expiration,
			HttpOnly: true,
			Secure:   true,
		})

		t := template.Must(template.ParseFiles("welcome.html"))
		t.Execute(w, nil)



	}

	registerHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
		if err != nil{
			log.Fatalf("failed to hash: %s", err)
		}

		err = database.AddUser(username, bcrypt)
		if err != nil{
			log.Fatalf("failed to add user: %s", err)
		}
		
		t := template.Must(template.ParseFiles("login.html"))
		t.Execute(w, nil)
	}

	secretHandler := func(w http.ResponseWriter, req *http.Request){
		claims, _ := req.Context().Value("claims").(*models.CustomClaims)
	
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to the protected endpoint, %s!", claims.Username)


	}
	bookmarkHandler := func(w http.ResponseWriter, req *http.Request){

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    	fmt.Fprint(w, `<i hx-post="/handleBookmarks" hx-trigger="click" hx-swap="outerHTML" class="far fa-bookmark text-blue-500  hover:text-white hover:bg-transparent cursor-pointer"></i> `)


	}
	checkBookmarkHandler := func(w http.ResponseWriter, req *http.Request){
		author := req.PostFormValue("author")
		title := req.PostFormValue("title")
		fmt.Println(author +" | "+ title)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    	fmt.Fprint(w, `<i hx-post="/handleBookmarks" hx-trigger="click" hx-swap="outerHTML" class="far fa-bookmark text-white  hover:text-blue  cursor-pointer"></i> `)
	}



	http.HandleFunc("/",MainPageHandler)
	http.HandleFunc("/entry",loginPageHandler)
	http.HandleFunc("/registerForm", registerPageHandler )	
	http.HandleFunc("/login",  middlewares.VerifyLogin(loginHandler) )
	http.HandleFunc("/register", middlewares.VerifyUser(registerHandler) )
	http.HandleFunc("/bookmarks", middlewares.VerifyJWT(secretHandler))
	http.HandleFunc("/handleBookmarks", middlewares.VerifyBookmarks(bookmarkHandler))
	http.HandleFunc("/checkBookmarks", checkBookmarkHandler )	

	log.Fatal(http.ListenAndServe(":8000",nil))
}