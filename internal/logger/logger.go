package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func LogsInit() {
	log = logrus.New()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/log.txt",
		MaxSize:    20,
		MaxAge:     3,
		MaxBackups: 5,
	})

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func LogRequest(method, path string) {
	log.Infof("[%s] %s", method, path)
}

func LogError(err error) {
	if err != nil {
		log.Errorf("Error: %v", err)
	}
}
