package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterTodosRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /todos", handlers.CreateTodo)
}
