package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/challenge/pkg/domain"
)

type UserRepository struct {
	Db sql.DB
}

func NewUserRepository(db sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (repo UserRepository) CreateUser(ctx context.Context, usr domain.User) int64 {
	insertUserSql := `INSERT INTO user(username, password) VALUES (?, ?)`

	statement, _ := repo.Db.Prepare(insertUserSql)

	result, _ := statement.Exec(usr.Username, usr.Password)

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	return id
}
