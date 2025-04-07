package errorhandler

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api/handlers"
	"github.com/wrtgvr/todoapi/internal/validation"
	rep "github.com/wrtgvr/todoapi/repository"
)

var errorsToStatusMap = map[error]int{
	// Bad Request
	handlers.ErrInvalidJSON:           http.StatusBadRequest,
	handlers.ErrInvalidUserID:         http.StatusBadRequest,
	handlers.ErrUserIdRequired:        http.StatusBadRequest,
	handlers.ErrTodoTitleRequired:     http.StatusBadRequest,
	validation.ErrPasswordTooShort:    http.StatusBadRequest,
	validation.ErrPasswordTooLong:     http.StatusBadRequest,
	validation.ErrPasswordRequired:    http.StatusBadRequest,
	validation.ErrUsernameTooShort:    http.StatusBadRequest,
	validation.ErrUsernameTooLong:     http.StatusBadRequest,
	validation.ErrUsernameRequired:    http.StatusBadRequest,
	validation.ErrDisallowedCharacter: http.StatusBadRequest,
	// Not Found
	rep.ErrUserNotFound: http.StatusNotFound,
	rep.ErrTodoNotFound: http.StatusNotFound,
}
