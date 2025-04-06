package router

import (
	"net/http"
	"strings"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterUsersRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request) {
		trimmedPath := strings.TrimPrefix(r.URL.Path, "/users/")
		parts := strings.Split(trimmedPath, "/")

		if len(parts) == 1 {
			handler.GetUser(w, r)
		} else if len(parts) == 2 {
			switch parts[1] {
			case "todos":
				handler.GetUserTodos(w, r)
			}
		} else {
			http.NotFound(w, r)
		}
	})
	mux.HandleFunc("DELETE /users/", handler.DeleteUser)
	mux.HandleFunc("PATCH /users/", handler.UpdateUser)
}
