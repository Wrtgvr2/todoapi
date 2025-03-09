package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log logrus.Logger

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/favicon.ico" {
			log.Infof("[%s] %s", r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	})
}

func LogsInit() {
	log := logrus.New()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/log.txt",
		MaxSize:    20,
		MaxAge:     3,
		MaxBackups: 5,
	})

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}
