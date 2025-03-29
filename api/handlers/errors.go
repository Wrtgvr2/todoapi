package handlers

import (
	"errors"
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
)

var (
	ErrInvalidJSON = errors.New("invalid json")
)

func HandleInternalError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
