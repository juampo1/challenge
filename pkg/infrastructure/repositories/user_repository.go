package repositories

import (
	"context"

	"github.com/challenge/pkg/domain"
)

type UserRepository struct {
	Db string
}

func NewUserRepository(db string) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (repo UserRepository) CreateUser(ctx context.Context, usr domain.User) (id int64) {
	return 1
}
