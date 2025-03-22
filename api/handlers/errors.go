package handlers

import (
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
)

func HandleInternalError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
