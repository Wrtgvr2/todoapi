package handlers

import rep "github.com/wrtgvr/todoapi/repository"

type Handler struct {
	UserRepo rep.UserRepo
	TodoRepo rep.TodoRepo
}
