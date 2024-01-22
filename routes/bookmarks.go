package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/williamjPriest/HTMXGO/controllers"
)

func LoadBookmarksHandler(w http.ResponseWriter, req *http.Request){
	loadedBookmarks := controllers.LoadBookmarkController(w,req)
	w.WriteHeader(http.StatusOK)
	t := template.Must(template.ParseGlob("templates/bookmarks.html"))
	t.Execute(w, loadedBookmarks)


}
func BookmarkHandler(w http.ResponseWriter, req *http.Request){
	added := controllers.BookmarkController(w,req)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if(added){
		fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
	}else{
		fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)

	}
	
}

func CheckBookmarkHandler(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)
}
