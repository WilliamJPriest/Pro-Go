package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"github.com/williamjPriest/HTMXGO/database"
)




type userData struct{
	Username string
	Password string
}

type ArticlesData struct{
	Articles []ArticleData `json:"articles"`
}


type ArticleData struct{
	Author string `json:"author"`
	Title string  `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	UrlToImage string `json:"urlToImage"`
}

type CustomClaims struct {
	Username string `json:"User"`
	jwt.StandardClaims
}


var SecretKey = []byte("SecretYouShouldHide")
func main(){
	// database.Create()
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	ApiKey := os.Getenv("API_KEY")
	
  
	MainPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("index.html"))
		res, err := http.Get(ApiKey)
		if err != nil{
			fmt.Println(err)
		}
		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		var responseObject ArticlesData

		json.Unmarshal(responseData, &responseObject)
		fmt.Println(responseObject.Articles[0].Author)


		t.Execute(w, responseObject)
	}	

	loginPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("login.html"))
		t.Execute(w, nil)
	}	

	registerPageHandler := func(w http.ResponseWriter, req *http.Request){
		t := template.Must(template.ParseFiles("register.html"))
		t.Execute(w, nil)
		

	}

	loginHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")


		token := jwt.New(jwt.SigningMethodHS256)
		expiration := time.Now().Add(10 * time.Minute)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = expiration.Unix()
		claims["authorized"] = true
		claims["user"] = username
	
		tokenString, err := token.SignedString(SecretKey)
		if err != nil {
			return 
		}
		
		fmt.Println(tokenString)
	
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expiration,
			HttpOnly: true,
			Secure:   true,
		})


	}

	registerHandler := func(w http.ResponseWriter, req *http.Request){
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		bcrypt,err := bcrypt.GenerateFromPassword([]byte(password), 5  )
		if err != nil{
			log.Fatalf("failed to hash: %s", err)
		}
		res, err := database.AddUser(username,bcrypt) 
		if err != nil{
			log.Fatalf("User couldn't be added to the database: %s", err)
		}
		fmt.Println(res)

	}

	secretHandler := func(w http.ResponseWriter, req *http.Request){
		claims, _ := req.Context().Value("claims").(*CustomClaims)
	
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to the protected endpoint, %s!", claims.Username)


	}



	http.HandleFunc("/",MainPageHandler)
	http.HandleFunc("/loginPage",loginPageHandler)
	http.HandleFunc("/registerPage", registerPageHandler )	
	http.HandleFunc("/login", loginHandler )
	http.HandleFunc("/register", registerHandler )
	http.HandleFunc("/secretData", verifyJWT(secretHandler))
	//convert to SPA possibily - articles

	log.Fatal(http.ListenAndServe(":8000",nil))
}
// try this with the generateJWT token 
func verifyJWT(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "No token provided")
			return
		}

		JWTstr := cookie.Value

		token, err := jwt.ParseWithClaims(JWTstr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Failed to parse token")
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

			ctx := context.WithValue(req.Context(), "claims", claims)
			req = req.WithContext(ctx)
			endpointHandler(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
		}
	})
}