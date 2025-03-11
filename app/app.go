package app

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/router"
	"github.com/wrtgvr/todoapi/internal/logger"
	"github.com/wrtgvr/todoapi/repository"
)

type App struct {
	Router http.Handler
}

func InitApp() (*App, error) {
	logger.LogsInit()

	errDb := repository.OpenDatabase()
	if errDb != nil {
		return nil, errDb
	}
	router := router.NewRouter()

	return &App{
		Router: router,
	}, nil
}

func CloseApp() {
	repository.CloseDatabase()
}
