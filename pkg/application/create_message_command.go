package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/challenge/pkg/domain"
)

const CreateMessageCommandName = "CreateMessageCommand"

type CreateMessageCommand struct {
	Sender      int64
	Recipient   int64
	ContentType string
	Text        string
}

type CreateMessageCommandHandler struct {
	MessageRepository MessageRepository
}

func (cmd CreateMessageCommand) GetName() string {
	return CreateMessageCommandName
}

func CreateMessageHandler(messageRepository MessageRepository) CreateMessageCommandHandler {
	return CreateMessageCommandHandler{
		MessageRepository: messageRepository,
	}
}

func (cm CreateMessageCommandHandler) Handle(ctx context.Context, cmd Command) (int64, time.Time, error) {
	msg, ok := cmd.(CreateMessageCommand)

	if !ok {
		fmt.Println("Wrong Command")
		return 0, time.Now(), errors.New("wrong command")
	}

	message := domain.NewMessage(domain.NewContent(msg.ContentType, msg.Text), time.Now(), msg.Sender, msg.Recipient)

	id, createdAt, err := cm.MessageRepository.CreateMessage(ctx, message)

	if err != nil {
		return 0, time.Now(), errors.New("something went wrong while storing the message")
	}

	return id, createdAt, nil
}
