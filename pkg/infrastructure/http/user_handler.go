package http

import "net/http"

type User struct {
	username string `json: username`
	password string `json: password`
}

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
