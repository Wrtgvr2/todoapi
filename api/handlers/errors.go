package handlers

import (
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
)

type error interface {
	Error() string
}

type ErrInvalidJSON struct{}

func (e ErrInvalidJSON) Error() string {
	return "Invalid JSON"
}

func HandleInternalError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
