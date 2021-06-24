package http

import (
	"net/http"

	"github.com/challenge/pkg/domain"
)

type Message struct {
	senderId       int64 `json: senderId`
	recipientId    int64 `json: recipientId`
	domain.Content `json: content`
}

func CreateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
