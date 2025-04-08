package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wrtgvr/todoapi/internal/validation"
	"github.com/wrtgvr/todoapi/models"
)

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	var bodyData models.UpdateTodoData

	err = DecodeBody(r.Body, &bodyData)
	if HandleError(w, err) {
		return
	}

	existingTodo, err := h.TodoRepo.GetTodo(id)
	if HandleError(w, err) {
		return
	}

	validatedData, err := ValidateAndPrepareUpdateTodoData(&bodyData, existingTodo)
	if HandleError(w, err) {
		return
	}

	todo, err := h.TodoRepo.UpdateTodo(id, validatedData)
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var bodyData models.CreateTodoData

	err := DecodeBody(r.Body, &bodyData)
	if HandleError(w, err) {
		return
	}
	_, err = h.UserRepo.GetUserById(*bodyData.User_ID)
	if HandleError(w, err) {
		return
	}

	err = validation.ValidateCreateTodoData(&bodyData)
	if HandleError(w, err) {
		return
	}

	_, err = h.UserRepo.GetUserById(*bodyData.User_ID)
	if HandleError(w, err) {
		return
	}

	createdTodo, err := h.TodoRepo.CreateToDo(&bodyData)
	if HandleError(w, err) {
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	err = h.TodoRepo.DeleteTodo(id)
	if HandleError(w, err) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	todo, err := h.TodoRepo.GetTodo(id)
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.TodoRepo.GetTodos()
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(todos)
}
