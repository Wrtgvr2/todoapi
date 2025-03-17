package models

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
