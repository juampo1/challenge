package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/challenge/pkg/domain"
)

const CreateUserCommandName = "CreateUserCommand"

type CreateUserCommand struct {
	Username string
	Password string
}

type CreateUserCommandHandler struct {
	UserRepository UserRepository
}

func (cmd CreateUserCommand) GetName() string {
	return CreateUserCommandName
}

func CreateUserHandler(userRepository UserRepository) CreateUserCommandHandler {
	return CreateUserCommandHandler{
		UserRepository: userRepository,
	}
}

func (cu CreateUserCommandHandler) Handle(ctx context.Context, cmd Command) (int64, error) {
	usr, ok := cmd.(CreateUserCommand)

	if !ok {
		fmt.Println("Wrong Command")
		return 0, errors.New("wrong command")
	}

	user := domain.NewUser(usr.Username, usr.Password)

	id := cu.UserRepository.CreateUser(ctx, user)

	return id, nil
}
