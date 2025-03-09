package handlers

import (
	"encoding/json"
	"net/http"
	//rep "github.com/wrtgvr/todoapi/repository"
)

type User struct {
	Username string   `json:"username"`
	Todos    []string `json:"todos"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := []User{ //! Remove before connection to DB
		{Username: "Wrtgvr", Todos: []string{"1", "2", "10", "41"}},
		{Username: "TestUser", Todos: []string{}},
	}

	json.NewEncoder(w).Encode(res)
}
