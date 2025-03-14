package models

type UserResponse struct {
	ID       uint64
	Username string
}

type UserRequest struct {
	Username string
	Password string
}
