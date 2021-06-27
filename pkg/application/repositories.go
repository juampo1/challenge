package application

import (
	"context"
	"time"

	"github.com/challenge/pkg/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, usr domain.User) int64
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, msg domain.Message) (int64, time.Time, error)
}
