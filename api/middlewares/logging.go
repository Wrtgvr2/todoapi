package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/favicon.ico" {
			log.Printf("[%s] %s", r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}
