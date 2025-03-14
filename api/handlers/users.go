package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/wrtgvr/todoapi/internal/logger"
	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var newUserData models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&newUserData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if newUserData.Username == "" || newUserData.Password == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if len(newUserData.Username) < 6 {
		http.Error(w, "Username must be at least 6 characters length", http.StatusBadRequest)
		return
	}
	if len(newUserData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters length", http.StatusBadRequest)
		return
	}

	user, err := rep.GetUserByUsername(newUserData.Username)
	if err != nil {
		logger.LogError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if user != nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	createdUser, err := rep.CreateUser(newUserData)
	if err != nil {
		logger.LogError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json, jsonErr := json.Marshal(createdUser)
	if jsonErr != nil {
		logger.LogError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		logger.LogError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := rep.GetUsers()
	if err != nil {
		logger.LogError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
