package handlers

import (
	"errors"
)

var (
	ErrInvalidJSON       = errors.New("invalid json")
	ErrInvalidUserID     = errors.New("invalid user id")
	ErrUserIdRequired    = errors.New("user id is required")
	ErrTodoTitleRequired = errors.New("title can't be empty")
)
