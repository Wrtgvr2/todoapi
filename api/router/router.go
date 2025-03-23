package router

import (
	"net/http"

	mws "github.com/wrtgvr/todoapi/api/middlewares"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	RegisterUsersRoutes(mux)
	RegisterTodosRoutes(mux)

	wrappedMux := mws.ChainMiddlewares(mux,
		mws.LoggingMiddleware,
		mws.SetHeaderJSON,
	)

	return wrappedMux
}
