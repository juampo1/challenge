package main

import (
	"log"
	"net/http"
	"time"

	"github.com/challenge/pkg/application"
	db "github.com/challenge/pkg/database"
	"github.com/challenge/pkg/infrastructure/auth"
	"github.com/challenge/pkg/infrastructure/controller"
	"github.com/challenge/pkg/infrastructure/repositories"
	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()

	//Handlers
	userHandler := controller.UserHandler{
		Cmd: application.CreateUserHandler(userRepository),
	}

	loginHandler := controller.LoginHandler{
		Qry:     application.CreateGetUserByUsernameQueryHandler(userRepository),
		JwtAuth: jwtAuth,
	}

	messageHandler := controller.MessageHandler{
		CreateMsgCmd:   application.CreateMessageHandler(messageRepository),
		GetMessagesQry: application.CreateGetMessagesQueryHandler(messageRepository),
	}

	//Users
	r.Post(UsersEndpoint, userHandler.CreateUser())

	//Login
	r.Post(LoginEndpoint, loginHandler.Login())

	//Message
	r.Post(MessagesEndpoint, auth.ValidateUser(messageHandler.CreateMessage(), key))
	r.Get(MessagesEndpoint, auth.ValidateUser(messageHandler.GetMessages(), key))

	//Health
	r.Post(CheckEndpoint, controller.Health())

	// Start server
	log.Println("Server started at port " + ServerPort)
	http.ListenAndServe(ServerPort, r)
	log.Fatal(http.ListenAndServe(ServerPort, nil))
}
