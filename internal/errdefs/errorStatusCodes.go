package errdefs

import (
	"net/http"
)

var errorsToStatusMap = map[error]int{
	// Bad Request
	ErrInvalidBody:          http.StatusBadRequest,
	ErrInvalidUserID:        http.StatusBadRequest,
	ErrUserIdRequired:       http.StatusBadRequest,
	ErrTodoTitleRequired:    http.StatusBadRequest,
	ErrPasswordTooShort:     http.StatusBadRequest,
	ErrPasswordTooLong:      http.StatusBadRequest,
	ErrPasswordRequired:     http.StatusBadRequest,
	ErrUsernameTooShort:     http.StatusBadRequest,
	ErrUsernameTooLong:      http.StatusBadRequest,
	ErrUsernameRequired:     http.StatusBadRequest,
	ErrDisallowedCharacters: http.StatusBadRequest,
	ErrTodoUserIdRequired:   http.StatusBadRequest,
	// Not Found
	ErrUserNotFound: http.StatusNotFound,
	ErrTodoNotFound: http.StatusNotFound,
	// Conflict
	ErrUsernameTaken: http.StatusConflict,
}
