package middlewares

import (
	"net/http"
	"github.com/williamjPriest/HTMXGO/database"
)

func VerifyBookmarks(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		title := req.PostFormValue("Title")
		database.CheckBookMarks(title)
		endpointHandler(w, req)
	})
}