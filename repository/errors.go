package repository

import (
	"fmt"
)

type error interface {
	Error() string
}

type ErrUserNotFound struct {
	RequestedID uint64
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("User with ID: [%d] not found.", e.RequestedID)
}
