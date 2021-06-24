package application

import "context"

type Command interface {
	Handle(ctx context.Context, cmd Command) error
}
