package validation

import (
	"regexp"

	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

var allowedPasswordCharset = regexp.MustCompile(`^[a-zA-Z0-9,.\\/<>{}"'?!$@#()%+=_-]+$`)
var allowerUsernameCharset = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var passwordMinLength = 8
var passwordMaxLength = 60
var usernameMinLength = 6
var usernameMaxLength = 24

func ValidatePassword(password string) error {
	if len(password) < passwordMinLength {
		return errdefs.ErrPasswordTooShort
	}
	if len(password) > passwordMaxLength {
		return errdefs.ErrPasswordTooLong
	}
	if !allowedPasswordCharset.MatchString(password) {
		return errdefs.ErrDisallowedCharacters
	}

	return nil
}

func ValidateUsername(username string) error {
	if len(username) < usernameMinLength {
		return errdefs.ErrUsernameTooShort
	}
	if len(username) > usernameMaxLength {
		return errdefs.ErrUsernameTooLong
	}
	if !allowerUsernameCharset.MatchString(username) {
		return errdefs.ErrDisallowedCharacters
	}

	return nil
}

func ValidateCreateUserRequest(userReq *models.UserRequest) error {
	if userReq.Username == nil {
		return errdefs.ErrUsernameRequired
	} else if err := ValidateUsername(*userReq.Username); err != nil {
		return err
	}
	if userReq.DisplayUsername != nil {
		if err := ValidateUsername(*userReq.DisplayUsername); err != nil {
			return err
		}
	}
	if userReq.Password == nil {
		return errdefs.ErrPasswordRequired
	} else if err := ValidatePassword(*userReq.Password); err != nil {
		return err
	}

	return nil
}
