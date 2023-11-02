package middlewares

import "net/http"

func VertifyLogin(endpointHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

	})
}