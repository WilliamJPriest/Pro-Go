package routes

import (
	"net/http"
	"text/template"
	"github.com/williamjPriest/HTMXGO/controllers"
)

func SearchHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/search.html"))	
	responseObject := controllers.SearchController(w,req)

	_, err := req.Cookie("token")
	if err != nil {
		t.Execute(w, responseObject)
	}else{
		t.Execute(w, responseObject)
	}
}