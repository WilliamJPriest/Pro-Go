package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/williamjPriest/HTMXGO/models"
)

func VerifyJWT(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("token")
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<a href="/entry" class="text-white no-underline"> You need to login brah</a> `)
			return
		}

		JWTstr := cookie.Value

		token, err := jwt.ParseWithClaims(JWTstr, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return models.SecretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Failed to parse token")
			return
		}

		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {

			ctx := context.WithValue(req.Context(), "claims", claims)
			req = req.WithContext(ctx)
			endpointHandler(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
		}
	})
}