package router

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
)

func RegisterTodosRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("GET /todos", handler.GetTodos)
	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("GET /todos/", handler.GetTodo)
	mux.HandleFunc("PATCH /todos/", handler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/", handler.DeleteTodo)
}
