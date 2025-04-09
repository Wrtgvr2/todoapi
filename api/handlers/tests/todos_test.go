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

func TestGetTodos(t *testing.T) {
	req := httptest.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()

	handler.GetTodos(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := []models.Todo{
		testTodoData,
	}

	var response []models.Todo
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestGetTodo_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/todos/1", nil)
	rec := httptest.NewRecorder()

	handler.GetTodo(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := testTodoData

	var response models.Todo
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestGetTodo_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/todos/2", nil)
	rec := httptest.NewRecorder()

	handler.GetTodo(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetTodo_BadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/todos/err", nil)
	rec := httptest.NewRecorder()

	handler.GetTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// DELETE
func TestDeleteTodo_Success(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/todos/1", nil)
	rec := httptest.NewRecorder()

	handler.DeleteTodo(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteTodo_NotFound(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/todos/2", nil)
	rec := httptest.NewRecorder()

	handler.DeleteTodo(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteTodo_BadRequest(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/todos/err", nil)
	rec := httptest.NewRecorder()

	handler.DeleteTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// PATCH
func TestUpdateTodo_Success(t *testing.T) {
	body, err := json.Marshal(testTodoUpdateData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/todos/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateTodo(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := testTodoData

	var response models.Todo
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestUpdateTodo_NotFound(t *testing.T) {
	body, err := json.Marshal(testTodoUpdateData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/todos/2", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateTodo(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateTodo_NoBody(t *testing.T) {
	req := httptest.NewRequest("PATCH", "/todos/2", nil)
	rec := httptest.NewRecorder()

	handler.UpdateTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateTodo_InvalidTitle(t *testing.T) {
	todoUpdateData := testTodoUpdateData
	todoUpdateData.Title = &testTodoTitle_BadReq
	body, err := json.Marshal(todoUpdateData)
	assert.NoError(t, err)

	req := httptest.NewRequest("PATCH", "/todos/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// POST
func TestCreateTodo_Success(t *testing.T) {
	todoData := testTodoCreateData
	body, err := json.Marshal(todoData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/todo", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateTodo(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedBody := testTodoData

	var response models.Todo
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedBody, response)
}

func TestCreateTodo_NoBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/todo", nil)
	rec := httptest.NewRecorder()

	handler.CreateTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateTodo_InvalidTitle(t *testing.T) {
	todoData := testTodoCreateData
	todoData.Title = &testTodoTitle_BadReq

	body, err := json.Marshal(todoData)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/todo", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateTodo(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
