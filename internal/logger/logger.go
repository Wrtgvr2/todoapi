package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func LogsInit() {
	log = logrus.New()

	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/%s-log.txt", date)

	for i := 1; fileExists(filename); i++ {
		filename = fmt.Sprintf("logs/%s-log-%d.txt", date, i)
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
