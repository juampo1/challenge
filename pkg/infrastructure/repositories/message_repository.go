package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/challenge/pkg/domain"
)

type MessageRepository struct {
	Db sql.DB
}

func NewMessageRepository(db sql.DB) *MessageRepository {
	return &MessageRepository{
		Db: db,
	}
}

func (repo MessageRepository) CreateMessage(ctx context.Context, msg domain.Message) (int64, time.Time, error) {

	return 1, time.Now(), nil
	// insertUserSql := `INSERT INTO user(username, password) VALUES (?, ?)`

	// statement, _ := repo.Db.Prepare(insertUserSql)

	// result, _ := statement.Exec(usr.Username, usr.Password)

	// id, err := result.LastInsertId()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// return id
}
