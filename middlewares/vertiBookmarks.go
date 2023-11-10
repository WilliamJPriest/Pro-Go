package middlewares

import (
	"fmt"
	"net/http"


	"github.com/williamjPriest/HTMXGO/utils"
	"github.com/williamjPriest/HTMXGO/database"
)

func VerifyBookmarks(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		title := req.PostFormValue("Title")
		username, err := utils.CheckUsername(req)
		if err != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<i hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML" class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer"></i> `)
			return
		}
		res := database.CheckBookMarks(title,username)
		if res != nil{
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<i hx-post="/handleBookmarks" hx-target="this" hx-trigger="click" hx-swap="outerHTML" class="far fa-bookmark text-white  hover:text-blue-500  cursor-pointer"></i> `)
			return
		}


	
		endpointHandler(w, req)
	})
}