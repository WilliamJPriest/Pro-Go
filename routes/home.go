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

func MainPageHandler(w http.ResponseWriter, req *http.Request){
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env file")
	}
	ApiKey := os.Getenv("API_KEY")
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	t := template.Must(template.ParseGlob("templates/index.html"))
	res, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&category=technology&"+ApiKey)
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