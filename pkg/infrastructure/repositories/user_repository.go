package repositories

import (
	"context"
	"database/sql"
	"errors"
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

func (repo UserRepository) GetUserByUsername(ctx context.Context, usr string) (domain.User, error) {
	var username string
	var password string

	getUserByUsernameSql := `SELECT username, password FROM user where username = ?`

	row := repo.Db.QueryRow(getUserByUsernameSql, usr)

	switch err := row.Scan(&username, &password); err {
	case sql.ErrNoRows:
		fmt.Println("username does not exists")
		return domain.User{}, errors.New("username does not exists")
	case nil:
		return domain.NewUser(username, password), nil
	default:
		panic(err)
	}
}
