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
	key := []byte("your-256-bit-secret")
	db, _ := db.Connect()
	jwtAuth := auth.NewJWTAuth("HS256", key, time.Hour)
	userRepository := repositories.NewUserRepository(*db)
	messageRepository := repositories.NewMessageRepository(*db)

	handlers := []httpx.ApiHandler{
		{
			Method:      "POST",
			Uri:         CheckEndpoint,
			HandlerFunc: httpx.Health(),
		},
		{
			Method:      "POST",
			Uri:         UsersEndpoint,
			HandlerFunc: httpx.CreateUser(application.CreateUserHandler(userRepository)),
		},
		{
			Method:      "POST",
			Uri:         LoginEndpoint,
			HandlerFunc: httpx.Login(application.CreateGetUserByUsernameQueryHandler(userRepository), jwtAuth),
		},
		{
			Method:      "POST",
			Uri:         MessagesEndpoint,
			HandlerFunc: httpx.CreateMessage(application.CreateMessageHandler(messageRepository), key),
		},
		{
			Method:      "GET",
			Uri:         MessagesEndpoint,
			HandlerFunc: httpx.GetMessages(application.CreateGetMessagesQueryHandler(messageRepository), key),
		},
	}

	httpHandler := httpx.SetUpHandlers(handlers...)

	// Start server
	log.Println("Server started at port " + ServerPort)
	http.ListenAndServe(ServerPort, httpHandler)
	log.Fatal(http.ListenAndServe(ServerPort, nil))
}
