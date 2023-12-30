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
	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/middlewares"
	"github.com/williamjPriest/HTMXGO/models"
	"github.com/williamjPriest/HTMXGO/utils"
	"golang.org/x/crypto/bcrypt"
)


func main(){
	err2 := godotenv.Load()
	if err2 != nil {
	  log.Fatal("Error loading .env file")
	}

	secretCode := os.Getenv("SECRET_CODE")
	var SecretKey = []byte(secretCode)

	apiKey := os.Getenv("API_KEY")
	
  
	MainPageHandler := func(w http.ResponseWriter, req *http.Request){
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		t := template.Must(template.ParseGlob("templates/index.html"))
		res, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&category=technology&"+apiKey)
		if err != nil{
			fmt.Println(err)
		}
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}


		var responseObject models.ArticlesData

		json.Unmarshal(responseData, &responseObject)

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
		t := template.Must(template.ParseGlob("templates/login.html"))
		t.Execute(w, nil)
	}	

	

	loginHandler := func(w http.ResponseWriter, req *http.Request){
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

	guestLoginHandler := func(w http.ResponseWriter, req *http.Request){
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

	logoutHandler := func(w http.ResponseWriter, req *http.Request){
		http.SetCookie(w, &http.Cookie{
			Name: "token",
			Expires: time.Now(),

		})
		w.Header().Set("Hx-Refresh", "true")
		

	}

	registerPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseGlob("templates/register.html"))
		t.Execute(w, nil)
		

	}

	registerHandler := func(w http.ResponseWriter, req *http.Request){
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

	loadBookmarksHandler := func(w http.ResponseWriter, req *http.Request){
		claims, _ := req.Context().Value("claims").(*models.CustomClaims)
		bookmark, err := database.GetBookMarks(claims.Username)
		if err != nil{
			fmt.Println("no bookmarks %w", err)
			return 
		}

		var bookmarks models.BookmarksData
		bookmarks.Username=claims.Username
		bookmarks.Bookmarks = bookmark
		
			
		
		w.WriteHeader(http.StatusOK)
		t := template.Must(template.ParseGlob("templates/bookmarks.html"))
		t.Execute(w, bookmarks)


	}
	bookmarkHandler := func(w http.ResponseWriter, req *http.Request){
		author := req.PostFormValue("Author")
		title := req.PostFormValue("Title")
		desc := req.PostFormValue("Description")
		url := req.PostFormValue("Url")
		urltoimage := req.PostFormValue("UrlToImage")
		username, err := utils.CheckUsername(req)
		if err != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
			return
		}

		if res := database.CheckBookMarks(title, username); res != nil{
			if err := database.AddBookMarks(author,title,desc,url,urltoimage,username); err != nil{
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
    			fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
    		fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)
			return
		}
		if del := database.RemovedBookMarks(title, username); del != nil{
			fmt.Println(del)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    	fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i> `)
		

	}
	checkBookmarkHandler := func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
    	fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)}
	
	searchHandler := func(w http.ResponseWriter, req *http.Request){		
		searchRes := req.PostFormValue("searchRes")
		t := template.Must(template.ParseGlob("templates/search.html"))
		res, err := http.Get("https://newsapi.org/v2/everything?q="+searchRes+"&language=en&"+apiKey)
		if err != nil{
			fmt.Println(err)
		}
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		var responseObject models.ArticlesData

		json.Unmarshal(responseData, &responseObject)

		_, err = req.Cookie("token")
		if err != nil {
			t.Execute(w, responseObject)
		}else{
			t.Execute(w, responseObject)
		}
	}
	





	http.HandleFunc("/",MainPageHandler)
	http.HandleFunc("/entry",loginPageHandler)
	http.HandleFunc("/login",  middlewares.VerifyLogin(loginHandler) )
	http.HandleFunc("/guestLogin",guestLoginHandler)
	http.HandleFunc("/logout",logoutHandler)	
	http.HandleFunc("/register", middlewares.VerifyUser(registerHandler) )
	http.HandleFunc("/registerForm", registerPageHandler )	
	http.HandleFunc("/bookmarks", middlewares.VerifyJWT(loadBookmarksHandler))
	http.HandleFunc("/handleBookmarks", middlewares.VerifyJWT(bookmarkHandler))
	http.HandleFunc("/checkBookmarks", middlewares.VerifyBookmarks(checkBookmarkHandler ))
	http.HandleFunc("/search", searchHandler)
	


	log.Fatal(http.ListenAndServe(":8000",nil))
}