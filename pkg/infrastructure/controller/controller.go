package controller

import (
	"github.com/challenge/pkg/application"
	"github.com/challenge/pkg/infrastructure/auth"
)

// Handler provides the interface to handle different requests
type UserHandler struct {
	Cmd application.CreateUserCommandHandler
}

type LoginHandler struct {
	Qry     application.GetUserByUsernameQueryHandler
	JwtAuth auth.JWTAuth
}

type MessageHandler struct {
	CreateMsgCmd   application.CreateMessageCommandHandler
	GetMessagesQry application.GetMessagesQueryHandler
}
