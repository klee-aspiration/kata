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
