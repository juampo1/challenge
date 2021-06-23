package controller

import (
	"net/http"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

// SendMessage send a message from one user to another
func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Send a New Message
	helpers.RespondJSON(w, domain.Message{})
}

// GetMessages get the messages from the logged user to a recipient
func (h Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages
	helpers.RespondJSON(w, []*domain.Message{{}})
}
