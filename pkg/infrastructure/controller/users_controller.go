package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/application"
	"github.com/challenge/pkg/helpers"
)

type user struct {
	Username string `json: username`
	Password string `json: password`
}

// CreateUser creates a new user
func (userHandler UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u user

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Something went wrong while parsing the user from request body"))
			return
		}

		createUserCmd := application.CreateUserCommand{
			Username: u.Username,
			Password: u.Password,
		}

		id, err := userHandler.Cmd.Handle(r.Context(), createUserCmd)

		if err != nil {
			httpError, _ := err.(helpers.HttpError)
			http.Error(w, httpError.Message, httpError.Code)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(id)
	}
}
