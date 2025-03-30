package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrtgvr/todoapi/api/handlers"
	"github.com/wrtgvr/todoapi/models"
)

var handler = handlers.Handler{
	UserRepo: MockUserRepo{},
}

// GET
func TestGetUsers(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	handler.GetUsers(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/1", nil)
	rec := httptest.NewRecorder()

	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/999", nil)
	rec := httptest.NewRecorder()

	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// DELETE
func TestDeleteUser(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteUser_NotFound(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/users/123123", nil)
	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// POST
func TestCreateUser(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateUser_InvalidUsername(t *testing.T) {
	badUsername := "qwe"
	userData := models.UserRequest{
		Username: &badUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateUser_InvalidPassword(t *testing.T) {
	badPassword := "qwe"
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &badPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// PATCH
func TestUpdateUser(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUser_InvalidID(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/131", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateUser_InvalidUsername(t *testing.T) {
	badUsername := "qwe"
	userData := models.UserRequest{
		Username: &badUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateUser_InvalidPassword(t *testing.T) {
	badPassword := "qwe"
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &badPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
