package api

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
	mws "github.com/wrtgvr/todoapi/api/middlewares"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", handlers.GetUsers)

	wrappedMux := mws.ChainMiddlewares(mux,
		mws.LoggingMiddleware,
	)

	return wrappedMux
}
