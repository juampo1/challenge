package application

import (
	"context"
)

type CreateUserCommand struct {
	username string
	password string
}

type CreateUserCommandHandler struct {
}

func (cu createUserCmd) Handle(ctx context.Context, cmd Command) error {

}
