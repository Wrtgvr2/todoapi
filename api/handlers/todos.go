package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var updateDataTodo models.UpdateTodoData

	err = json.NewDecoder(r.Body).Decode(&updateDataTodo)
	if err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	if strings.ReplaceAll(*updateDataTodo.Title, " ", "") == "" {
		http.Error(w, ErrTodoTitleRequired.Error(), http.StatusBadRequest)
		return
	}

	existingTodo, err := h.TodoRepo.GetTodo(id)
	if err != nil {
		if errors.Is(err, rep.ErrTodoNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	if updateDataTodo.Completed == nil {
		updateDataTodo.Completed = existingTodo.Completed
	}
	if updateDataTodo.Description == nil {
		updateDataTodo.Description = existingTodo.Description
	}
	if updateDataTodo.Title == nil {
		updateDataTodo.Title = &existingTodo.Title
	}

	todo, err := h.TodoRepo.UpdateTodo(id, &updateDataTodo)
	if err != nil { // Don't check for ErrTodoNotFound cuz checks it earlier
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.TodoRepo.DeleteTodo(id)
	if err != nil {
		if errors.Is(err, rep.ErrTodoNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	todo, err := h.TodoRepo.GetTodo(id)
	if err != nil {
		if errors.Is(err, rep.ErrTodoNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.TodoRepo.GetTodos()
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var bodyData *models.CreateTodoData

	if err := json.NewDecoder(r.Body).Decode(&bodyData); err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.UserRepo.GetUserById(bodyData.User_ID); err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	if strings.ReplaceAll(bodyData.Title, " ", "") == "" {
		http.Error(w, ErrTodoTitleRequired.Error(), http.StatusBadRequest)
		return
	}

	createdTodo, err := h.TodoRepo.CreateToDo(bodyData)
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}
