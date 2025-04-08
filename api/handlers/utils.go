package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/internal/logger"
	"github.com/wrtgvr/todoapi/internal/validation"
	"github.com/wrtgvr/todoapi/models"
	"golang.org/x/crypto/bcrypt"
)

func GetIdFromUrl(urlPath string) (uint64, error) {
	r := regexp.MustCompile(`^\/(?:users|todos)/(\d+)(:?/todos)?`)
	var id uint64

	matches := r.FindStringSubmatch(urlPath)
	if len(matches) > 1 {
		var err error
		id, err = strconv.ParseUint(matches[1], 10, 64)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errdefs.ErrInvalidUserID
	}

	return id, nil
}

func HandleInternalError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	logger.LogError(err)
}

func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashBytes), nil
}

func ValidateAndPrepareCreateUserRequest(userReq *models.UserRequest) (*models.UserRequest, error) {
	preparedData := *userReq
	if err := validation.ValidateCreateUserRequest(&preparedData); err != nil {
		return nil, err
	}
	if userReq.DisplayUsername == nil {
		preparedData.DisplayUsername = userReq.Username
	}
	return &preparedData, nil
}

func ValidateAndPrepareUserData(userReq *models.UserRequest, existingUser *models.User) (*models.User, error) {
	var newUserData models.User

	if userReq.Username != nil {
		if err := validation.ValidateUsername(*userReq.Username); err != nil {
			return nil, err
		}
		newUserData.Username = *userReq.Username
	} else if existingUser != nil {
		newUserData.Username = existingUser.Username
	} else {
		return nil, errdefs.ErrUsernameRequired
	}

	if userReq.DisplayUsername != nil {
		if err := validation.ValidateUsername(*userReq.DisplayUsername); err != nil {
			return nil, err
		}
		newUserData.DisplayUsername = *userReq.DisplayUsername
	} else if existingUser != nil {
		newUserData.DisplayUsername = existingUser.DisplayUsername
	} else {
		newUserData.DisplayUsername = newUserData.Username
	}

	if userReq.Password != nil {
		if err := validation.ValidatePassword(*userReq.Password); err != nil {
			return nil, err
		}
		var err error
		hashedPassword, err := HashPassword(*userReq.Password)
		if err != nil {
			if err == bcrypt.ErrPasswordTooLong {
				return nil, errdefs.ErrPasswordTooShort
			}
			return nil, err
		}
		newUserData.Password = hashedPassword

	} else if existingUser != nil {
		newUserData.Password = existingUser.Password
	} else {
		return nil, errdefs.ErrPasswordRequired
	}

	if existingUser != nil {
		newUserData.ID = existingUser.ID
	}

	return &newUserData, nil
}

func ConvertUserToUserRequest(user *models.User) *models.UserRequest {
	var userReq = models.UserRequest{
		Username:        &user.Username,
		Password:        &user.Password,
		DisplayUsername: &user.DisplayUsername,
	}

	return &userReq
}

func HandleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	fmt.Println(err.Error())
	statusCode := errdefs.GetErrorStatusCode(err)
	if statusCode == http.StatusInternalServerError {
		HandleInternalError(w, err)
	} else {
		http.Error(w, err.Error(), statusCode)
	}

	return true
}

func ValidateAndPrepareUpdateTodoData(todoData *models.UpdateTodoData, existingTodo *models.Todo) (*models.UpdateTodoData, error) {
	var validatedData models.UpdateTodoData

	if todoData.Title != nil {
		err := validation.ValidateTitle(*todoData.Title)
		if err != nil {
			return nil, err
		}
		validatedData.Title = todoData.Title
	} else if existingTodo != nil {
		validatedData.Title = &existingTodo.Title
	} else {
		return nil, errdefs.ErrTodoTitleRequired
	}

	if todoData.Description != nil {
		validatedData.Description = todoData.Description
	} else if existingTodo != nil {
		validatedData.Description = &existingTodo.Description
	} else {
		emptyDecription := ""
		validatedData.Description = &emptyDecription
	}

	if todoData.Completed != nil {
		validatedData.Completed = todoData.Completed
	} else if existingTodo != nil {
		validatedData.Completed = &existingTodo.Completed
	} else {
		notCompleted := false
		validatedData.Completed = &notCompleted
	}

	return &validatedData, nil
}

func DecodeBody(body io.ReadCloser, v any) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		if err == io.EOF {
			return errdefs.ErrInvalidBody
		}
		return err
	}

	return nil
}
