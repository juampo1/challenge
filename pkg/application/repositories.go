package application

import (
	"context"

	"github.com/challenge/pkg/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, usr domain.User) int64
}
