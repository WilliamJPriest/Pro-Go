package main

import(
	"net/http"
	// "fmt"
	"io"
	"log"
	// "github.com/gin-gonic/gin"
)

func main(){
	heyHandler := func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w, "<h1> Hey There</h1> \n")
	}

	http.HandleFunc("/gay",heyHandler)
	log.Fatal(http.ListenAndServe(":3000",nil))
}