package middlewares

import (
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/favicon.ico" {
			logger.LogRequest(r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}
