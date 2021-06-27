package domain

import "time"

type Message struct {
	Sender    int64
	Recipient int64
	CreatedAt time.Time
	Content
}

func NewMessage(content Content, createdAt time.Time, sender, recipient int64) Message {
	return Message{
		Sender:    sender,
		Recipient: recipient,
		CreatedAt: createdAt,
		Content:   NewContent(content.ContentType, content.Text),
	}
}
