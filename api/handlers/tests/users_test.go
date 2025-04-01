package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrtgvr/todoapi/models"
)

// GET
func TestGetUsers(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	handler.GetUsers(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := []models.UserResponse{
		{ID: 1, Username: TestUsername},
		{ID: 2, Username: TestUsername},
	}

	var response []models.UserResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestGetUser_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/1", nil)
	rec := httptest.NewRecorder()

	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := models.UserResponse{
		ID:       TestUserID,
		Username: TestUsername,
	}
	var response models.UserResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestGetUser_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/999", nil)
	rec := httptest.NewRecorder()

	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetUser_BadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/users/error", nil)
	rec := httptest.NewRecorder()

	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// DELETE
func TestDeleteUser_Success(t *testing.T) {
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

func TestDeleteUser_BadRequest(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/users/errerr", nil)
	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// POST
func TestCreateUser_Success(t *testing.T) {
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

	expectedBody := models.UserResponse{
		ID:       TestUserID,
		Username: TestUsername,
	}
	var response models.UserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestCreateUser_InvalidUsername(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername_BadReq,
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
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword_BadReq,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// PATCH
func TestUpdateUser_Success(t *testing.T) {
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

	expectedBody := models.UserResponse{
		ID:       TestUserID,
		Username: TestUsername,
	}

	var response models.UserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestUpdateUser_NotFound(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/2", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateUser_InvalidUsername(t *testing.T) {
	userData := models.UserRequest{
		Username: &TestUsername_BadReq,
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
	userData := models.UserRequest{
		Username: &TestUsername,
		Password: &TestPassword_BadReq,
	}
	body, err := json.Marshal(userData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/users/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
