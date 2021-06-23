package controller

import (
	"net/http"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Create a New User
	helpers.RespondJSON(w, domain.User{})
}
