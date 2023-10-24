package main

import(
	"net/http"
	"html/template"
	"fmt"
	// "io"
	"log"
	// "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
var ApiKey string = "https://newsapi.org/v2/top-headlines?country=us&category=tech&apiKey=2f4376c9e22f40c7aa18a7e783a566d3"

type userData struct{
	Username string
	Password string
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

func main(){
	MainPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		// response, err := http.Get(ApiKey)
		// if err != nil{
		// 	log.Fatalf("response issue: %s", err)
		// }
		// fmt.Println(response)
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
		Password := req.PostFormValue("password")
		if Username != MUserName {
			log.Fatalf("This name didn't match: %s", Username)
  
		}
		err := bcrypt.CompareHashAndPassword([]byte(MPassword) , []byte(Password))
	    if err != nil{
			log.Fatalf("didn't match: %s", err)
		}
		fmt.Println("Yo")
		
		

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