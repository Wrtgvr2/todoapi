package errdefs

import (
	"errors"
)

var (
	ErrInvalidBody          = errors.New("invalid body")
	ErrInvalidUserID        = errors.New("invalid user id")
	ErrUserIdRequired       = errors.New("user id is required")
	ErrTodoTitleRequired    = errors.New("title can't be empty")
	ErrUserNotFound         = errors.New("user not found")
	ErrTodoNotFound         = errors.New("todo not found")
	ErrPasswordTooShort     = errors.New("password must be at least 8 characters long")
	ErrPasswordTooLong      = errors.New("password cannot be longer than 60 characters")
	ErrUsernameTooShort     = errors.New("username must be at least 6 characters long")
	ErrUsernameTooLong      = errors.New("username cannot be longer than 24 characters")
	ErrDisallowedCharacters = errors.New("disallowed characters used")
	ErrUsernameRequired     = errors.New("username is required")
	ErrPasswordRequired     = errors.New("password is required")
	ErrTodoUserIdRequired   = errors.New("user id required for todo")
	ErrUsernameTaken        = errors.New("username already taken")
)
