package handlers_test

import (
	"time"

	"github.com/wrtgvr/todoapi/models"
)

// Users
var (
	TestUsername        string = "testUser"
	TestPassword        string = "testPassword"
	TestUserID          uint64 = 1
	TestUsername_BadReq string = "err"
	TestPassword_BadReq string = "err"
)

// Todos
var (
	TestTodoID           uint64    = 1
	TestTodoUserID       uint64    = 1
	TestTodoTitle        string    = "testtitle"
	TestTodoTitle_BadReq string    = " "
	TestTodoDescription  string    = "test"
	TestTodoCompleted    bool      = false
	TestTodoCreatedAt    time.Time = time.Date(2008, time.September, 16, 15, 19, 26, 0, time.UTC)
)

var TestTodoUpdateData = models.UpdateTodoData{
	Title:       &TestTodoTitle,
	Description: &TestTodoDescription,
	Completed:   &TestTodoCompleted,
}

var TestTodoCreateData = models.CreateTodoData{
	User_ID:     TestUserID,
	Title:       TestTodoTitle,
	Description: &TestTodoDescription,
}

var TestTodoData = models.Todo{
	ID:          TestTodoID,
	User_ID:     TestTodoUserID,
	Title:       TestTodoTitle,
	Description: &TestTodoDescription,
	Completed:   &TestTodoCompleted,
	Created_At:  TestTodoCreatedAt,
}
