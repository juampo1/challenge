package application

import "context"

type Command interface {
	GetName() string
}

type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) (int64, error)
}
