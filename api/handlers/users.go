package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
	"golang.org/x/crypto/bcrypt"
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

	var requestUserData models.UserRequest
	err = json.NewDecoder(r.Body).Decode(&requestUserData)
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

	if requestUserData.Username == nil {
		updatedUserData.Username = existingUser.Username
	} else {
		updatedUserData.Username = *requestUserData.Username
	}

	if requestUserData.Password == nil {
		updatedUserData.Password = existingUser.Password
	} else {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(*requestUserData.Password), bcrypt.DefaultCost)
		if err != nil {
			if err == bcrypt.ErrPasswordTooLong {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			HandleInternalError(w, err)
			return
		}
		updatedUserData.Password = string(hashBytes)
	}

	if requestUserData.Username != nil && len(updatedUserData.Username) < 6 {
		http.Error(w, "Username must be at least 6 characters length", http.StatusBadRequest)
		return
	}
	if requestUserData.Password != nil && len(updatedUserData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters length", http.StatusBadRequest)
		return
	}

	userWithSameUsername, err := h.UserRepo.GetUserByUsername(updatedUserData.Username)
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
	var requestUserData models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&requestUserData); err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	if *requestUserData.Username == "" || len(*requestUserData.Username) < 6 {
		http.Error(w, "Username must be at least 6 characters length", http.StatusBadRequest)
		return
	}
	if *requestUserData.Password == "" || len(*requestUserData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters length", http.StatusBadRequest)
		return
	}

	_, err := h.UserRepo.GetUserByUsername(*requestUserData.Username)
	if err != nil {
		if !errors.Is(err, rep.ErrUserNotFound) {
			HandleInternalError(w, err)
			return
		}
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(*requestUserData.Password), bcrypt.DefaultCost)
	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		HandleInternalError(w, err)
		return
	}
	hashedPassword := string(hashBytes)

	var newUserData models.UserRequest
	newUserData.Password = &hashedPassword
	newUserData.Username = requestUserData.Username

	fmt.Println(hashedPassword)

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
