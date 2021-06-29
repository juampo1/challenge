package application

import (
	"context"

	"github.com/challenge/pkg/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, usr domain.User) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, msg domain.Message) (int64, error)
	GetMessages(ctx context.Context, recipient int64, start int64) ([]domain.Message, error)
}
