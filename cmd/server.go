package main

import (
	"log"
	"net/http"

	httpx "github.com/challenge/pkg/infrastructure/http"
)

const (
	ServerPort       = ":8080"
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)

func main() {

	handlers := []httpx.ApiHandler{
		{
			Method:      "POST",
			Uri:         "/health",
			HandlerFunc: httpx.Health(),
		},
	}

	httpHandler := httpx.SetUpHandlers(handlers...)

	// h := controller.Handler{}

	// // Configure endpoints
	// // Health
	// http.HandleFunc(CheckEndpoint, func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	h.Check(w, r)
	// })

	// // Users
	// http.HandleFunc(UsersEndpoint, func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	h.CreateUser(w, r)
	// })

	// // Auth
	// http.HandleFunc(LoginEndpoint, func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	h.Login(w, r)
	// })

	// // Messages
	// http.HandleFunc(MessagesEndpoint, auth.ValidateUser(func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		h.GetMessages(w, r)
	// 	case http.MethodPost:
	// 		h.SendMessage(w, r)
	// 	default:
	// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// }))

	// Start server
	log.Println("Server started at port " + ServerPort)
	http.ListenAndServe(ServerPort, httpHandler)
	log.Fatal(http.ListenAndServe(ServerPort, nil))
}
