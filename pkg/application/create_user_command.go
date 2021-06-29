package application

import (
	"context"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
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
		return 0, helpers.NewInternalServerError("Wrong command")
	}

	if usr.Username == "" || usr.Password == "" {
		return 0, helpers.NewBadRequestError("Username/Password cannot be empty")
	}

	user := domain.NewUser(usr.Username, usr.Password)

	id, err := cu.UserRepository.CreateUser(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
