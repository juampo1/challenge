package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/challenge/pkg/application"
	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

type Content struct {
	ContentType string `json: "contentType"`
	Text        string `json: "text"`
}

type Message struct {
	Id        int64     `json: id`
	Timestamp time.Time `json: timestamp`
	Sender    int64     `json: "sender"`
	Recipient int64     `json: "recipient"`
	Content   Content   `json: "content"`
}

func (messageHandler MessageHandler) CreateMessage() http.HandlerFunc {
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

		id, err := messageHandler.CreateMsgCmd.Handle(r.Context(), createMsgCommand)

		if err != nil {
			httpError, _ := err.(helpers.HttpError)
			http.Error(w, httpError.Message, httpError.Code)
			return
		}

		json.NewEncoder(w).Encode(struct {
			Id        int64
			Timestamp time.Time
		}{
			Id:        id,
			Timestamp: time.Now()})
	}
}

func (messageHandler MessageHandler) GetMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipient, _ := strconv.ParseInt(r.URL.Query().Get("recipient"), 10, 64)
		start, _ := strconv.ParseInt(r.URL.Query().Get("start"), 10, 64)

		getMessagesQuery := application.GetMessagesQuery{
			Recipient: recipient,
			Start:     start,
		}

		messages, err := messageHandler.GetMessagesQry.Handle(r.Context(), getMessagesQuery)

		if err != nil {
			httpError, _ := err.(helpers.HttpError)
			http.Error(w, httpError.Message, httpError.Code)
			return
		}

		json.NewEncoder(w).Encode(struct {
			Messages []domain.Message
		}{Messages: messages})
	}
}
