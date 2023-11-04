package controllers

import (
	"log"
	"net/http"

	"github.com/williamjPriest/HTMXGO/middlewares"
	"github.com/williamjPriest/HTMXGO/routes"
)

func Controllers(){
	http.HandleFunc("/", routes.MainPageHandler)
	http.HandleFunc("/entry",routes.loginPageHandler)
	http.HandleFunc("/registerForm", routes.registerPageHandler )	
	http.HandleFunc("/login",  middlewares.VerifyLogin(routes.loginHandler) )
	http.HandleFunc("/register", middlewares.VerifyUser(routes.registerHandler) )
	http.HandleFunc("/secretData", middlewares.VerifyJWT(routes.secretHandler))

	log.Fatal(http.ListenAndServe(":8000",nil))
}
