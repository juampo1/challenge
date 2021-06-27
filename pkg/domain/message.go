package domain

import "time"

type Message struct {
	sender    int64
	recipient int64
	createdAt time.Time
	Content
}

func NewMessage(content Content, createdAt time.Time, sender, recipient int64) Message {
	return Message{
		sender:    sender,
		recipient: recipient,
		createdAt: createdAt,
		Content:   NewContent(content.contentType, content.text),
	}
}
