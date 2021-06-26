package main

import (
	"log"
	"net/http"

	"github.com/challenge/pkg/application"
	db "github.com/challenge/pkg/database"
	httpx "github.com/challenge/pkg/infrastructure/http"
	"github.com/challenge/pkg/infrastructure/repositories"
)

const (
	ServerPort       = ":8080"
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)

func main() {

	db, _ := db.Connect()
	userRepository := repositories.NewUserRepository(*db)

	handlers := []httpx.ApiHandler{
		{
			Method:      "POST",
			Uri:         "/check",
			HandlerFunc: httpx.Health(),
		},
		{
			Method:      "POST",
			Uri:         "/users",
			HandlerFunc: httpx.CreateUser(application.CreateUserHandler(userRepository)),
		},
		{
			Method:      "POST",
			Uri:         "/check",
			HandlerFunc: httpx.Health(),
		},
		{
			Method:      "POST",
			Uri:         "/login",
			HandlerFunc: httpx.Login(application.CreateGetUserByUsernameQueryHandler(userRepository)),
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
