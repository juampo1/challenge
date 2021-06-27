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
	insertMessageSql := `INSERT INTO message(sender_id, recipient_id, created_at) VALUES (?, ?, CURRENT_TIMESTAMP)`
	insertContentSql := `INSERT INTO content(message_id, content_type, text) VALUES (?, ?, ?)`

	insertMessageStatement, _ := repo.Db.Prepare(insertMessageSql)
	insertContentStatement, _ := repo.Db.Prepare(insertContentSql)

	insertMeesageResult, _ := insertMessageStatement.Exec(msg.Sender, msg.Recipient)

	messageId, _ := insertMeesageResult.LastInsertId()

	_, _ = insertContentStatement.Exec(messageId, msg.Content.ContentType, msg.Content.Text)

	return messageId, time.Now(), nil
}
