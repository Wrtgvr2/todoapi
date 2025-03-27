package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var logsPath = "./logs"
var logFile *os.File
var maxLogFiles = 5

func InitLogs() {
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/%s-log", date)

	maxSuffix := getLogFileMaxSuffix()
	if maxSuffix >= 0 {
		filename = fmt.Sprintf("%s-%d", filename, maxSuffix+1)
	}

	filename = fmt.Sprintf("%s.%s", filename, "txt")

	var err error
	logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	for countLogFiles() > maxLogFiles {
		os.Remove(getOldestLogFilePath())
	}
}

func CloseLogs() {
	err := logFile.Close()
	if err != nil {
		panic(err)
	}
}

func LogMessage(msg string) {
	writeLog(msg, "Message")
}

func LogRequest(method, path string) {
	text := fmt.Sprintf("[%s] %s", method, path)
	writeLog(text, "Request")
}

func LogError(err error) {
	if err != nil {
		writeLog(err.Error(), "Error")
	}
}

func getLogFileMaxSuffix() int {
	maxSuffix := -1
	r := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}-log(?:-(\d+))?\.txt$`)

	files, err := os.ReadDir(logsPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		match := r.FindStringSubmatch(file.Name())
		if len(match) > 1 && match[1] != "" {
			suffix, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			if suffix > maxSuffix {
				maxSuffix = suffix
			}
		}
	}

	if maxSuffix == -1 {
		maxSuffix = 0
	}

	return maxSuffix
}

func getOldestLogFilePath() string {
	files, err := os.ReadDir(logsPath)
	if err != nil {
		panic(err)
	}

	var oldestFilePath string
	var oldestTime time.Time

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", logsPath, file.Name())

		info, err := os.Stat(filePath)
		if err != nil {
			continue
		}

		if oldestFilePath == "" || info.ModTime().Before(oldestTime) {
			oldestFilePath = filePath
			oldestTime = info.ModTime()
		}
	}

	return oldestFilePath
}

func countLogFiles() int {
	files, err := os.ReadDir(logsPath)
	if err != nil {
		panic(err)
	}

	return len(files)
}

func writeLog(text, logType string) {
	logMsg := map[string]string{
		"type": logType,
		"text": text,
		"date": time.Now().Format("Jan 2006-01-02 15-04-05"),
	}

	json, err := json.Marshal(logMsg)
	if err != nil {
		panic(err)
	}

	_, err = logFile.Write(json)
	if err != nil {
		panic(err)
	}
}
