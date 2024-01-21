package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/williamjPriest/HTMXGO/controllers"
	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/utils"
)

func LoadBookmarksHandler(w http.ResponseWriter, req *http.Request){
	loadedBookmarks := controllers.LoadBookmarkController(w,req)
	w.WriteHeader(http.StatusOK)
	t := template.Must(template.ParseGlob("templates/bookmarks.html"))
	t.Execute(w, loadedBookmarks)


}
func BookmarkHandler(w http.ResponseWriter, req *http.Request){
	author := req.PostFormValue("Author")
	title := req.PostFormValue("Title")
	desc := req.PostFormValue("Description")
	url := req.PostFormValue("Url")
	urltoimage := req.PostFormValue("UrlToImage")
	username, err := utils.CheckUsername(req)
	if err != nil{
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
		return
	}

	resChan := make(chan error)

	go func(){
		res:= database.CheckBookMarks(title, username)
		resChan <-res
	}()
	res := <-resChan
	if res != nil{
		if err := database.AddBookMarks(author,title,desc,url,urltoimage,username); err != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)
		return
	}
	if del := database.RemovedBookMarks(title, username); del != nil{
		fmt.Println(del)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i> `)
	

}
func CheckBookmarkHandler(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-blue-500  hover:text-white  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-white  hover:text-blue-500 cursor-pointer"></i>`)}
