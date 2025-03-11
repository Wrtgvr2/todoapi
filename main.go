package main

import (
	"net/http"

	"github.com/wrtgvr/todoapi/app"
	"github.com/wrtgvr/todoapi/internal/logger"
)

func main() {
	App, err := app.InitApp()
	if err != nil {
		logger.LogError(err)
		return
	}

	http.ListenAndServe(":8080", App.Router)
}
