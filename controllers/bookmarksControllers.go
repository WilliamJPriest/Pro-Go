package controllers

import (
	"fmt"
	"net/http"

	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/models"
)

func LoadBookmarkController(w http.ResponseWriter, req *http.Request)(loadedBookmarks models.BookmarksData){
	claims, _ := req.Context().Value("claims").(*models.CustomClaims)
	bookmark, err := database.GetBookMarks(claims.Username)
	if err != nil{
		fmt.Println("no bookmarks %w", err)
		return 
	}

	var bookmarks models.BookmarksData
	bookmarks.Username=claims.Username
	bookmarks.Bookmarks = bookmark
	return bookmarks

}