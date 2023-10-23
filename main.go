package main

import(
	"net/http"
	"html/template"
	// "fmt"
	"io"
	"log"
	// "github.com/gin-gonic/gin"
)

func main(){
	heyHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(w, nil)
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w, "<h1>HI</h1>")
	}
	
	http.HandleFunc("/hey",heyHandler)
	http.HandleFunc("/btn", helloHandler )
	log.Fatal(http.ListenAndServe(":8000",nil))
}