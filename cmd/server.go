package main

import (
	"log"
	"net/http"
	"time"

	"github.com/challenge/pkg/application"
	db "github.com/challenge/pkg/database"
	"github.com/challenge/pkg/infrastructure/auth"
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
	jwtAuth := auth.NewJWTAuth("HS256", []byte("your-256-bit-secret"), time.Hour)
	userRepository := repositories.NewUserRepository(*db)
	messageRepository := repositories.NewMessageRepository(*db)

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
			Uri:         "/login",
			HandlerFunc: httpx.Login(application.CreateGetUserByUsernameQueryHandler(userRepository), jwtAuth),
		},
		{
			Method:      "POST",
			Uri:         "/messages",
			HandlerFunc: httpx.CreateMessage(application.CreateMessageHandler(messageRepository)),
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
