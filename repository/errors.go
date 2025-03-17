package repository

type error interface {
	Error() string
}

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "User not found."
}
