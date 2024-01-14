package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"


	"github.com/williamjPriest/HTMXGO/middlewares"
	"github.com/williamjPriest/HTMXGO/models"
	"github.com/williamjPriest/HTMXGO/routes"


)


func main(){


	ApiKey := os.Getenv("API_KEY")

  

	searchHandler := func(w http.ResponseWriter, req *http.Request){		
		searchRes := req.PostFormValue("searchRes")
		t := template.Must(template.ParseGlob("templates/search.html"))
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
	





	http.HandleFunc("/",routes.MainPageHandler)
	http.HandleFunc("/entry",routes.LoginPageHandler)
	http.HandleFunc("/login",  middlewares.VerifyLogin(routes.LoginHandler) )
	http.HandleFunc("/guestLogin",routes.GuestLoginHandler)
	http.HandleFunc("/logout",routes.LogoutHandler)	
	http.HandleFunc("/register", middlewares.VerifyUser(routes.RegisterHandler) )
	http.HandleFunc("/registerForm", routes.RegisterPageHandler )	
	http.HandleFunc("/bookmarks", middlewares.VerifyJWT(routes.LoadBookmarksHandler))
	http.HandleFunc("/handleBookmarks", middlewares.VerifyJWT(routes.BookmarkHandler))
	http.HandleFunc("/checkBookmarks", middlewares.VerifyBookmarks(routes.CheckBookmarkHandler ))
	http.HandleFunc("/search", searchHandler)

	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}


	log.Fatal(http.ListenAndServe("0.0.0.0:"+ port,nil))
}