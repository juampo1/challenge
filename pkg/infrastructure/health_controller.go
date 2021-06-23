package controller

import (
	"net/http"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

// Check returns the health of the service and DB
func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	// TODO: Check service health. Feel free to add any check you consider necessary
	helpers.RespondJSON(w, domain.Health{})
}
