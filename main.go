package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(port, App.Router); err != nil {
			log.Fatalf("Can't start app: %v\n", err)
		}
	}()

	<-stopCh

	app.CloseApp()
}
