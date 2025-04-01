package handlers_test

import "github.com/wrtgvr/todoapi/api/handlers"

var handler = handlers.Handler{
	UserRepo: MockUserRepo{},
	TodoRepo: MockTodoRepo{},
}
