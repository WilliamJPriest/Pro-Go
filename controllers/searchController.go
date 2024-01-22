package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/models"
)

func SearchController(w http.ResponseWriter, req *http.Request)(resObject models.ArticlesData){
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
	return responseObject
}