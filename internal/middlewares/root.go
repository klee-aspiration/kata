package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charliemcelfresh/kata/internal/config"
)

func EnforceAPIKataRequestContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != config.Constants["REQUIRED_API_KATA_REQUEST_CONTENT_TYPE"] {
			http.Error(w, fmt.Sprintf("Content-Type header must be %v",
				config.Constants["REQUIRED_API_KATA_REQUEST_CONTENT_TYPE"]),
				http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf(
			`timestamp: %v, request_path: %v`,
			time.RFC3339, r.RequestURI,
		))
		next.ServeHTTP(w, r)
	})
}

func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentTypeToSet := fmt.Sprintf("%v", config.Constants["REQUIRED_API_KATA_RESPONSE_CONTENT_TYPE"])
		w.Header().Set("Content-Type", contentTypeToSet)
		next.ServeHTTP(w, r)
	})
}
