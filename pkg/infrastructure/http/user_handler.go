package http

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Username string `json: username`
	Password string `json: password`
}

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u user

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Something went wrong while parsing the user from request body"))
			return
		}
	}
}
