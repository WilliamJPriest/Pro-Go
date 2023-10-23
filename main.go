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

var MPassword []byte

func main(){
	loginPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(w, nil)
	}	

	registerPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("register.html"))
		t.Execute(w, nil)

	}

	loginHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		fmt.Println(username)
		err := bcrypt.CompareHashAndPassword([]byte(MPassword) , []byte(password))
		fmt.Println(err == nil)
		// if err != nil{
		// 	log.Fatalf("didn't match: %s", err)
		// }
		
		

	}
	registerHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
		if err != nil{
			log.Fatalf("failed to hash: %s", err)
		}
		fmt.Println(username)
		MPassword=bcrypt

	}


	
	http.HandleFunc("/loginPage",loginPageHandler)
	http.HandleFunc("/registerPage", registerPageHandler )	
	http.HandleFunc("/login", loginHandler )
	http.HandleFunc("/register", registerHandler )	

	log.Fatal(http.ListenAndServe(":8000",nil))
}