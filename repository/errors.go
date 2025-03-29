package repository

import "errors"

var (
	ErrUserNotFound = errors.New("User not found.")
	ErrTodoNotFound = errors.New("Todo not found.")
)
