package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	err = h.UserRepo.DeleteUser(id)
	if HandleError(w, err) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	var requestUserData models.UserRequest
	err = DecodeBody(r.Body, &requestUserData)
	if HandleError(w, err) {
		return
	}

	actualUserData, err := h.UserRepo.GetFullUser(id)
	if HandleError(w, err) {
		return
	}

	validatedUserData, err := ValidateAndPrepareUserData(&requestUserData, actualUserData)
	if HandleError(w, err) {
		return
	}

	userWithSameUsername, err := h.UserRepo.GetUserByUsername(validatedUserData.Username)
	if err != nil && !errors.Is(err, errdefs.ErrUserNotFound) {
		HandleError(w, err)
		return
	}
	if userWithSameUsername != nil && userWithSameUsername.ID != validatedUserData.ID {
		HandleError(w, errdefs.ErrUsernameTaken)
		return
	}

	updatedUser, err := h.UserRepo.UpdateUser(validatedUserData)
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestUserData models.UserRequest

	err := DecodeBody(r.Body, &requestUserData)
	if HandleError(w, err) {
		return
	}

	preparedUserData, err := ValidateAndPrepareCreateUserRequest(&requestUserData)
	if HandleError(w, err) {
		return
	}

	userWithSameUsername, err := h.UserRepo.GetUserByUsername(*preparedUserData.Username)
	if err != nil && !errors.Is(err, errdefs.ErrUserNotFound) {
		HandleError(w, err)
		return
	}
	if userWithSameUsername != nil {
		HandleError(w, errdefs.ErrUsernameTaken)
		return
	}

	createdUser, err := h.UserRepo.CreateUser(preparedUserData)
	if HandleError(w, err) {
		return
	}

	json, jsonErr := json.Marshal(createdUser)
	if HandleError(w, jsonErr) {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	user, err := h.UserRepo.GetUserById(id)
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.UserRepo.GetUsers()
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetUserTodos(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if HandleError(w, err) {
		return
	}

	_, err = h.UserRepo.GetUserById(id)
	if HandleError(w, err) {
		return
	}

	todo, err := h.UserRepo.GetUserTodos(id)
	if HandleError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(todo)
}
