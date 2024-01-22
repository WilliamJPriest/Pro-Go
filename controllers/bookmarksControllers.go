package controllers

import (
	"fmt"
	"net/http"

	"github.com/williamjPriest/HTMXGO/database"
	"github.com/williamjPriest/HTMXGO/models"
	"github.com/williamjPriest/HTMXGO/utils"
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

func BookmarkController(w http.ResponseWriter, req *http.Request)(added bool){
	author := req.PostFormValue("Author")
	title := req.PostFormValue("Title")
	desc := req.PostFormValue("Description")
	url := req.PostFormValue("Url")
	urltoimage := req.PostFormValue("UrlToImage")
	username, err := utils.CheckUsername(req)
	if err != nil{
		return false
	}

	resChan := make(chan error)

	go func(){
		res:= database.CheckBookMarks(title, username)
		resChan <-res
	}()
	res := <-resChan			
	if res != nil{
		if add := database.AddBookMarks(author,title,desc,url,urltoimage,username); add != nil{
			return false
		}
		return true
	}
	database.RemovedBookMarks(title, username); 
	return false

}