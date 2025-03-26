package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = rep.DeleteUser(id)
	if err != nil {
		if notFoundErr, ok := err.(*rep.ErrUserNotFound); ok {
			http.Error(w, notFoundErr.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	var updateData models.UserRequest
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, ErrInvalidJSON{}.Error(), http.StatusBadRequest)
		return
	}

	existingUser, err := rep.GetFullUser(id)
	if err != nil {
		if notFoundErr, ok := err.(*rep.ErrUserNotFound); ok {
			http.Error(w, notFoundErr.Error(), http.StatusNotFound)
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

	userWithSameUsername, err := rep.GetUserByUsername(*updateData.Username)
	if err != nil {
		if _, ok := err.(*rep.ErrUserNotFound); !ok {
			HandleInternalError(w, err)
			return
		}
	}
	if userWithSameUsername != nil && userWithSameUsername.ID != existingUser.ID {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	updatedUser, err := rep.UpdateUser(&updatedUserData)
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserData models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&newUserData); err != nil {
		http.Error(w, ErrInvalidJSON{}.Error(), http.StatusBadRequest)
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

	_, err := rep.GetUserByUsername(*newUserData.Username)
	if err != nil {
		if _, ok := err.(*rep.ErrUserNotFound); !ok {
			HandleInternalError(w, err)
			return
		}
	}

	createdUser, err := rep.CreateUser(&newUserData)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	user, err := rep.GetUserById(id)
	if err != nil {
		if notFoundErr, ok := err.(*rep.ErrUserNotFound); ok {
			http.Error(w, notFoundErr.Error(), http.StatusNotFound)
			return
		}
		HandleInternalError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := rep.GetUsers()
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
