package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterUsersRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users/", handler.GetUser)
	mux.HandleFunc("DELETE /users/", handler.DeleteUser)
	mux.HandleFunc("PATCH /users/", handler.UpdateUser)
}
