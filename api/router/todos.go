package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterTodosRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /todos", handlers.GetTodos)
	mux.HandleFunc("GET /todos/", handlers.GetTodo)
	mux.HandleFunc("POST /todos", handlers.CreateTodo)
}
