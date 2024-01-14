package main

import (
	"log"
	"net/http"
	"os"

	"github.com/williamjPriest/HTMXGO/middlewares"

	"github.com/williamjPriest/HTMXGO/routes"
)


func main(){

	http.HandleFunc("/",routes.MainPageHandler)
	http.HandleFunc("/entry",routes.LoginPageHandler)
	http.HandleFunc("/login",  middlewares.VerifyLogin(routes.LoginHandler) )
	http.HandleFunc("/guestLogin",routes.GuestLoginHandler)
	http.HandleFunc("/logout",routes.LogoutHandler)	
	http.HandleFunc("/register", middlewares.VerifyUser(routes.RegisterHandler))
	http.HandleFunc("/registerForm", routes.RegisterPageHandler )	
	http.HandleFunc("/bookmarks", middlewares.VerifyJWT(routes.LoadBookmarksHandler))
	http.HandleFunc("/handleBookmarks", middlewares.VerifyJWT(routes.BookmarkHandler))
	http.HandleFunc("/checkBookmarks", middlewares.VerifyBookmarks(routes.CheckBookmarkHandler ))
	http.HandleFunc("/search", routes.SearchHandler)

	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}


	log.Fatal(http.ListenAndServe("0.0.0.0:"+ port,nil))
}