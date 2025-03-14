package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterUsersRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("POST /users", handlers.PostUser)
	mux.HandleFunc("GET /users/", handlers.GetUser)
}
