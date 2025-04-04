package main

import (
	"net/http"

	"github.com/wrtgvr/todoapi/app"
	"github.com/wrtgvr/todoapi/internal/logger"
)

const port = ":8080"

func main() {
	App, err := app.InitApp()
	if err != nil {
		logger.LogError(err)
		return
	}
	defer app.CloseApp()

	http.ListenAndServe(port, App.Router)
}
