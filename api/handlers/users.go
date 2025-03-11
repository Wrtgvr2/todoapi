package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
	rep "github.com/wrtgvr/todoapi/repository"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := rep.GetUsers()
	if err != nil {
		logger.LogError(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
