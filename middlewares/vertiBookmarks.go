package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/utils"
)

func VerifyBookmarks(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}
		title := req.PostFormValue("Title")
		username, err := utils.CheckUsername(req)

		if err != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500 hover:text-white  cursor-pointer"></i> `)
			return
		}
		res := database.CheckBookMarks(title,username)
		if res != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<div hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML"> <i class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer" ></i><i class="htmx-indicator far fa-bookmark text-blue-500  hover:text-white cursor-pointer"></i>`)
			return
		}


	
		endpointHandler(w, req)
	})
}