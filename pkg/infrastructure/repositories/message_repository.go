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

func (repo MessageRepository) GetMessages(ctx context.Context, recipient int64, start int64) ([]domain.Message, error) {
	messages := []domain.Message{}

	getMessagesSql := `SELECT m.recipient_id, m.sender_id, m.created_at, c.content_type, c.text
											From message m 
											INNER JOIN content c ON m.id = c.message_id
											WHERE m.recipient_id = ? AND m.id >= ?
											ORDER BY m.id ASC
											LIMIT 100
												`

	getMessageStatement, _ := repo.Db.Prepare(getMessagesSql)

	messageRows, _ := getMessageStatement.Query(recipient, start)

	for messageRows.Next() {
		var message domain.Message
		var content domain.Content

		messageRows.Scan(&message.Recipient, &message.Sender, &message.CreatedAt)
		messageRows.Scan(&content.ContentType, &content.Text)

		messages = append(messages, message)
	}

	return messages, nil
}
