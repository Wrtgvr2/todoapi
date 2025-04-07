package validation

import (
	"errors"
	"fmt"
)

var (
	ErrPasswordTooShort    = fmt.Errorf("password must be at least %d characters long", passwordMinLength)
	ErrPasswordTooLong     = fmt.Errorf("password cannot be longer than %d characters", passwordMaxLength)
	ErrUsernameTooShort    = fmt.Errorf("username must be at least %d characters long", usernameMinLength)
	ErrUsernameTooLong     = fmt.Errorf("username cannot be longer than %d characters", usernameMaxLength)
	ErrDisallowedCharacter = errors.New("input contains unsupported characters")
	ErrUsernameRequired    = errors.New("username is required")
	ErrPasswordRequired    = errors.New("password is required")
)
