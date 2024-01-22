package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/models"
)

func SearchHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/search.html"))	
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env file")
	}
	ApiKey := os.Getenv("API_KEY")	
	searchRes := req.PostFormValue("searchRes")
	
	res, err := http.Get("https://newsapi.org/v2/everything?q="+searchRes+"&language=en&"+ApiKey)
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