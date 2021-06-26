package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/challenge/pkg/domain"
)

const GetUserByUsernameQueryName = "GetUserByUsername"

type GetUserByUsernameQuery struct {
	Username string
	Password string
}

type GetUserByUsernameQueryHandler struct {
	UserRepository UserRepository
}

func (qry GetUserByUsernameQuery) GetQueryName() string {
	return GetUserByUsernameQueryName
}

func CreateGetUserByUsernameQueryHandler(userRepository UserRepository) GetUserByUsernameQueryHandler {
	return GetUserByUsernameQueryHandler{
		UserRepository: userRepository,
	}
}

func (qry GetUserByUsernameQueryHandler) Handle(ctx context.Context, query Query) (domain.User, error) {
	usr, ok := query.(GetUserByUsernameQuery)

	if !ok {
		fmt.Println("wrong query")
		return domain.User{}, errors.New("wrong query")
	}

	user, err := qry.UserRepository.GetUserByUsername(ctx, usr.Username)

	if err != nil {
		return domain.User{}, err
	}

	if usr.Username == user.Username && usr.Password == user.Password {
		return user, nil
	}

	return user, nil
}
