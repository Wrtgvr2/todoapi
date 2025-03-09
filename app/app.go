package app

import (
	"database/sql"
	"net/http"

	"github.com/wrtgvr/todoapi/api"
	"github.com/wrtgvr/todoapi/internal/logger"
	"github.com/wrtgvr/todoapi/repository"
)

type App struct {
	DB     *sql.DB
	Router http.Handler
}

func InitApp() (*App, error) {
	logger.LogsInit()

	db, errDb := repository.OpenDatabase()
	if errDb != nil {
		return nil, errDb
	}
	router := api.NewRouter()

	return &App{
		DB:     db,
		Router: router,
	}, nil
}

func CloseApp() {
	if err := repository.CloseDatabase(); err != nil {
		logger.LogError(err)
	}
}
