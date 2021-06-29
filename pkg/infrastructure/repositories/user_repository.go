package repositories

import (
	"context"
	"database/sql"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

type UserRepository struct {
	Db sql.DB
}

func NewUserRepository(db sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (repo UserRepository) CreateUser(ctx context.Context, usr domain.User) (int64, error) {
	insertUserSql := `INSERT INTO user(username, password) VALUES (?, ?)`

	statement, _ := repo.Db.Prepare(insertUserSql)

	//Password must be encrypted before storing it, never store it as plain text.
	result, err := statement.Exec(usr.Username, usr.Password)

	if err != nil {
		helpers.NewInternalServerError("something went wrog while storing user")
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (repo UserRepository) GetUserByUsername(ctx context.Context, usr string) (domain.User, error) {
	var id int64
	var username string
	var password string

	getUserByUsernameSql := `SELECT id, username, password FROM user where username = ?`

	row := repo.Db.QueryRow(getUserByUsernameSql, usr)

	switch err := row.Scan(&id, &username, &password); err {
	case sql.ErrNoRows:
		return domain.User{}, helpers.NewNotFoundError("User does not exist")
	case nil:
		return domain.User{
			Id:       id,
			Username: username,
			Password: password,
		}, nil
	default:
		return domain.User{}, helpers.NewInternalServerError("Something went wrong while retrieving the user from database")
	}
}
