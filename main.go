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

type userData struct{
	username string
	password string
}

type ArticleData struct{
	author string
	title string
	description string
	url string
	urlToImage string
}

var MUserName string
var MPassword []byte

func main(){
	MainPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		response, err := http.Get(ApiKey)
		if err != nil{
			log.Fatalf("response issue: %s", err)
		}
		fmt.Println(response)
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
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		if username != MUserName {
			log.Fatalf("This name didn't match: %s", username)
  
		}
		err := bcrypt.CompareHashAndPassword([]byte(MPassword) , []byte(password))
	    if err != nil{
			log.Fatalf("didn't match: %s", err)
		}
		fmt.Println("Yo")
		
		

	}
	registerHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
		if err != nil{
			log.Fatalf("failed to hash: %s", err)
		}
		MUserName= username
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