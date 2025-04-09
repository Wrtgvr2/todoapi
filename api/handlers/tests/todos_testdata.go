package handlers_test

import (
	"time"

	"github.com/wrtgvr/todoapi/models"
)

var (
	testTodoID           uint64    = 1
	testTodoUserID       uint64    = 1
	testTodoTitle        string    = "testtitle"
	testTodoTitle_BadReq string    = " "
	testTodoDescription  string    = "test"
	testTodoCompleted    bool      = false
	testTodoCreatedAt    time.Time = time.Date(2008, time.September, 16, 15, 19, 26, 0, time.UTC)
)

var testTodoUpdateData = models.UpdateTodoData{
	Title:       &testTodoTitle,
	Description: &testTodoDescription,
	Completed:   &testTodoCompleted,
}

var testTodoCreateData = models.CreateTodoData{
	User_ID:     &testUserID,
	Title:       &testTodoTitle,
	Description: testTodoDescription,
}

var testTodoData = models.Todo{
	ID:          testTodoID,
	User_ID:     testTodoUserID,
	Title:       testTodoTitle,
	Description: testTodoDescription,
	Completed:   testTodoCompleted,
	Created_At:  testTodoCreatedAt,
}
