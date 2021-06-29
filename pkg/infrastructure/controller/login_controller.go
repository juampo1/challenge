package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/application"
	"github.com/challenge/pkg/helpers"
)

func (h LoginHandler) Login() http.HandlerFunc {
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

		user, err := h.Qry.Handle(r.Context(), getUserByUsernameQuery)

		if err != nil {
			httpError, _ := err.(helpers.HttpError)
			http.Error(w, httpError.Message, httpError.Code)
			return
		}

		token, _ := h.JwtAuth.CreateToken(user.Id)

		json.NewEncoder(w).Encode(token)
	}
}
