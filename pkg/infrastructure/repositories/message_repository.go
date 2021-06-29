package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

type MessageRepository struct {
	Db sql.DB
}

func NewMessageRepository(db sql.DB) *MessageRepository {
	return &MessageRepository{
		Db: db,
	}
}

func (repo MessageRepository) CreateMessage(ctx context.Context, msg domain.Message) (int64, error) {
	insertMessageSql := `INSERT INTO message(sender_id, recipient_id, created_at) VALUES (?, ?, CURRENT_TIMESTAMP)`
	insertContentSql := `INSERT INTO content(message_id, content_type, text) VALUES (?, ?, ?)`

	insertMessageStatement, _ := repo.Db.Prepare(insertMessageSql)
	insertContentStatement, _ := repo.Db.Prepare(insertContentSql)

	insertMeesageResult, err := insertMessageStatement.Exec(msg.Sender, msg.Recipient)

	if err != nil {
		return 0, helpers.NewInternalServerError("Something went wrong while storing message")
	}

	messageId, _ := insertMeesageResult.LastInsertId()

	_, err = insertContentStatement.Exec(messageId, msg.Content.ContentType, msg.Content.Text)

	if err != nil {
		return 0, helpers.NewInternalServerError("Something went wrong while storing content")
	}

	insertMessageStatement.Close()
	insertContentStatement.Close()

	return messageId, nil
}

func (repo MessageRepository) GetMessages(ctx context.Context, recipient int64, start int64) ([]domain.Message, error) {
	messages := []domain.Message{}

	getMessagesSql := `SELECT m.recipient_id, m.sender_id, c.content_type, c.text, m.created_at
											From message m 
											INNER JOIN content c ON m.id = c.message_id
											WHERE m.recipient_id = ? AND m.id >= ?
											ORDER BY m.id ASC
											LIMIT 100`

	getMessageStatement, _ := repo.Db.Prepare(getMessagesSql)

	messageRows, err := getMessageStatement.Query(recipient, start)

	if err != nil {
		return nil, helpers.NewInternalServerError("Something went wrong while retrieving messages")
	}

	for messageRows.Next() {
		var recipient int64
		var sender int64
		var createdAt time.Time
		var contentType string
		var text string

		messageRows.Scan(&recipient, &sender, &contentType, &text, &createdAt)
		message := domain.NewMessage(domain.NewContent(contentType, text), createdAt, sender, recipient)
		messages = append(messages, message)
	}

	return messages, nil
}
