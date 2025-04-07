package validation

import (
	"regexp"
)

var allowedPasswordCharset = regexp.MustCompile(`^[a-zA-Z0-9,.\\/<>{}"'?!$@#()%+=_\-]+$`)
var allowerUsernameCharset = regexp.MustCompile(`^[a-zA-Z0-9_]$`)
var passwordMinLength = 8
var passwordMaxLength = 60
var usernameMinLength = 6
var usernameMaxLength = 24

func ValidatePassword(password string) error {
	if len(password) < passwordMinLength {
		return ErrPasswordTooShort
	}
	if len(password) > passwordMaxLength {
		return ErrPasswordTooLong
	}
	if !allowedPasswordCharset.MatchString(password) {
		return ErrDisallowedCharacter
	}

	return nil
}

func ValidateUsername(username string) error {
	if len(username) < usernameMinLength {
		return ErrUsernameTooShort
	}
	if len(username) > usernameMaxLength {
		return ErrPasswordTooLong
	}
	if !allowerUsernameCharset.MatchString(username) {
		return ErrDisallowedCharacter
	}

	return nil
}
