package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := rep.GetTodos()
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var bodyData models.CreateTodoData

	if err := json.NewDecoder(r.Body).Decode(&bodyData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if _, err := rep.GetUserById(bodyData.User_ID); err != nil {
		if notFoundErr, ok := err.(*rep.ErrUserNotFound); ok {
			http.Error(w, notFoundErr.Error(), http.StatusNotFound)
			return
		}
	}
	if strings.ReplaceAll(bodyData.Title, " ", "") == "" {
		http.Error(w, "Title can't be empty", http.StatusBadRequest)
		return
	}

	createdTodo, err := rep.CreateToDo(bodyData)
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}
