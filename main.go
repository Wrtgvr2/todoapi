package main

import (
	"net/http"

	"github.com/wrtgvr/todoapi/api"
)

func main() {
	r := api.NewRouter()

	http.ListenAndServe(":8080", r)
}
