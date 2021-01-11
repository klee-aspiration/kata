package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf(
			`timestamp: %v, request_path: %v`,
			time.RFC3339, r.RequestURI,
		))
		next.ServeHTTP(w, r)
	})
}

func EnforceContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Content-Type") != "application/kata.api-json" {
			response.WriteHeader(415)
		} else {
			handler.ServeHTTP(response, request)
		}
	})
}
