package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/challenge/pkg/application"
)

type Content struct {
	ContentType string `json: "contentType,omitempty"`
	Text        string `json: "text,omitempty"`
}
type Message struct {
	Id        int64     `json: id`
	Timestamp time.Time `json: timestamp`
	Sender    int64     `json: "sender,omitempty"`
	Recipient int64     `json: "recipient,omitempty"`
	Content   Content   `json: "content,omitempty"`
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
			Sender:      msg.Sender,
			Recipient:   msg.Recipient,
			ContentType: msg.Content.ContentType,
			Text:        msg.Content.Text,
		}

		id, timestamp, _ := cmd.Handle(r.Context(), createMsgCommand)

		json.NewEncoder(w).Encode(Message{Id: id, Timestamp: timestamp})
	}
}

func GetMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
