package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/challenge/pkg/application"
)

type Content struct {
	ContentType string `json: contentType`
	Text        string `json: text`
}

type Message struct {
	Sender    int64   `json: sender`
	Recipient int64   `json: recipient`
	Content   Content `json: content`
}

func CreateMessage(cmd application.CreateMessageCommandHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var msg Message

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Something went wrong while parsing the message from request body"))
			return
		}

		createMsgCommand := application.CreateMessageCommand{
			Sender:    msg.Sender,
			Recipient: msg.Recipient,
			CreatedAt: time.Now(),
			Content:   msg.Content,
		}

		id, _, _ := cmd.Handle(r.Context(), createMsgCommand)

		json.NewEncoder(w).Encode(id)
	}
}

func GetMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
