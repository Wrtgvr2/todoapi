package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
	mws "github.com/wrtgvr/todoapi/api/middlewares"
	rep "github.com/wrtgvr/todoapi/repository"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	userRepo := &rep.PostgresUserRepo{DB: rep.DB}
	todoRepo := &rep.PostgresTodoRepo{DB: rep.DB}

	handler := &handlers.Handler{
		UserRepo: userRepo,
		TodoRepo: todoRepo,
	}

	RegisterUsersRoutes(mux, handler)
	RegisterTodosRoutes(mux, handler)

	wrappedMux := mws.ChainMiddlewares(mux,
		mws.LoggingMiddleware,
		mws.SetHeaderJSON,
	)

	return wrappedMux
}
