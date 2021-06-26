package application

import (
	"context"

	"github.com/challenge/pkg/domain"
)

type Query interface {
	GetQueryName() string
}

type QueryHandler interface {
	Handle(ctx context.Context, query Query) (domain.User, error)
}
