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

func MainController(w http.ResponseWriter, req *http.Request)(resObject models.ArticlesData){
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env file")
	}
	ApiKey := os.Getenv("API_KEY")
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

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
	return responseObject
}

