package handlers

import (
	"errors"
	"net/http"

	"github.com/wrtgvr/todoapi/internal/logger"
)

var (
	ErrInvalidJSON       = errors.New("invalid json")
	ErrInvalidUserID     = errors.New("invalid user id")
	ErrUserIdRequired    = errors.New("user id is required")
	ErrTodoTitleRequired = errors.New("title can't be empty")
)

func HandleInternalError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
