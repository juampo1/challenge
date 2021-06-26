package http

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/application"
)

type user struct {
	Username string `json: username`
	Password string `json: password`
}

func CreateUser(cmd application.CreateUserCommandHandler) http.HandlerFunc {
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

		id, _ := cmd.Handle(r.Context(), createUserCmd)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(id)
	}
}
