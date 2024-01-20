package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/williamjPriest/HTMXGO/models"
)

func CheckUsername(req *http.Request) (string, error){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	secretCode := os.Getenv("SECRET_CODE")
	var SecretKey = []byte(secretCode)
	cookie, err := req.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	JWTstr := cookie.Value

	token, err := jwt.ParseWithClaims(JWTstr, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return "",fmt.Errorf("%w", err)
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims.Username, nil
	}
	return "", nil
}

