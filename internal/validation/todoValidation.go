package validation

import (
	"strings"

	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

func ValidateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return errdefs.ErrTodoTitleRequired
	}
	return nil
}

func ValidateCreateTodoData(todo *models.CreateTodoData) error {
	err := ValidateTitle(*todo.Title)
	if todo.Title == nil || err != nil {
		return errdefs.ErrTodoTitleRequired
	}
	if todo.User_ID == nil {
		return errdefs.ErrTodoUserIdRequired
	}

	return nil
}
