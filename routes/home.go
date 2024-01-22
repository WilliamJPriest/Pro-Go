package routes

import (
	"net/http"
	"text/template"
	"github.com/williamjPriest/HTMXGO/controllers"

)

func MainPageHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/index.html"))
	responseObject := controllers.MainController(w,req)

	_, err := req.Cookie("token")
	if err != nil {
		responseObject.IsLoggedIn = false
		t.Execute(w, responseObject)
	}else{
		responseObject.IsLoggedIn = true
		t.Execute(w, responseObject)
	}
	
	
}	