package http

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/application"
	"github.com/challenge/pkg/infrastructure/auth"
)

func Login(qry application.GetUserByUsernameQueryHandler, jwtAuth auth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u user

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Something went wrong while parsing the user from request body"))
			return
		}

		getUserByUsernameQuery := application.GetUserByUsernameQuery{
			Username: u.Username,
			Password: u.Password,
		}

		user, _ := qry.Handle(r.Context(), getUserByUsernameQuery)

		token, _ := jwtAuth.CreateToken(user.Id)

		json.NewEncoder(w).Encode(token)
	}
}
