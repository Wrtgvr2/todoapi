package handlers_test

import "github.com/wrtgvr/todoapi/api/handlers"

var Handler = handlers.Handler{
	UserRepo: MockUserRepo{},
	TodoRepo: MockTodoRepo{},
}
