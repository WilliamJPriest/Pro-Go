package routes

import (
	"net/http"
	"text/template"
	"time"
	"github.com/williamjPriest/HTMXGO/controllers"

)


func LoginPageHandler(w http.ResponseWriter, req *http.Request){
	t := template.Must(template.ParseGlob("templates/login.html"))
	t.Execute(w, nil)
}	

func LoginHandler(w http.ResponseWriter, req *http.Request){
	controllers.LoginController(w, req)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.ParseGlob("templates/welcome.html"))
	t.Execute(w,  req.PostFormValue("username"))
}

func GuestLoginHandler(w http.ResponseWriter, req *http.Request){
	controllers.GuestLoginController(w, req)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.ParseGlob("templates/welcome.html"))
	t.Execute(w, "Guest101")
}

func LogoutHandler(w http.ResponseWriter, req *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Expires: time.Now(),

	})
	w.Header().Set("HX-Redirect", "/entry")
}

func RegisterPageHandler(w http.ResponseWriter, req *http.Request){	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.ParseGlob("templates/register.html"))
	t.Execute(w, nil)

}

func RegisterHandler(w http.ResponseWriter, req *http.Request){
	controllers.RegisterController(w, req)
	t := template.Must(template.ParseGlob("templates/login.html"))
	t.Execute(w, nil)
}