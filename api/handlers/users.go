package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.UserRepo.DeleteUser(id)
	if err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var updateData models.UserRequest
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	existingUser, err := h.UserRepo.GetFullUser(id)
	if err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	var updatedUserData models.User

	updatedUserData.ID = id

	if updateData.Username == nil {
		updatedUserData.Username = existingUser.Username
	} else {
		updatedUserData.Username = *updateData.Username
	}

	if updateData.Password == nil {
		updatedUserData.Password = existingUser.Password
	} else {
		updatedUserData.Password = *updateData.Password
	}

	if updateData.Username != nil && len(updatedUserData.Username) < 6 {
		http.Error(w, "Username must be at least 6 characters length", http.StatusBadRequest)
		return
	}
	if updateData.Password != nil && len(updatedUserData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters length", http.StatusBadRequest)
		return
	}

	userWithSameUsername, err := h.UserRepo.GetUserByUsername(*updateData.Username)
	if err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	if userWithSameUsername != nil && userWithSameUsername.ID != existingUser.ID {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	updatedUser, err := h.UserRepo.UpdateUser(&updatedUserData)
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserData models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&newUserData); err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	if *newUserData.Username == "" || len(*newUserData.Username) < 6 {
		http.Error(w, "Username must be at least 6 characters length", http.StatusBadRequest)
		return
	}
	if *newUserData.Password == "" || len(*newUserData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters length", http.StatusBadRequest)
		return
	}

	_, err := h.UserRepo.GetUserByUsername(*newUserData.Username)
	if err != nil {
		if !errors.Is(err, rep.ErrUserNotFound) {
			HandleInternalError(w, err)
			return
		}
	}

	createdUser, err := h.UserRepo.CreateUser(&newUserData)
	if err != nil {
		HandleInternalError(w, err)
		return
	}
	json, jsonErr := json.Marshal(createdUser)
	if jsonErr != nil {
		HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.UserRepo.GetUserById(id)
	if err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.UserRepo.GetUsers()
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetUserTodos(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdFromUrl(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.UserRepo.GetUserTodos(id)
	if err != nil {
		if errors.Is(err, rep.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
