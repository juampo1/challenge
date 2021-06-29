package application

import (
	"context"
	"errors"
	"time"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
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

func (cm CreateMessageCommandHandler) Handle(ctx context.Context, cmd Command) (int64, error) {
	msg, ok := cmd.(CreateMessageCommand)

	if !ok {
		return 0, helpers.NewInternalServerError("Wrong command")
	}

	message := domain.NewMessage(domain.NewContent(msg.ContentType, msg.Text), time.Now(), msg.Sender, msg.Recipient)

	id, err := cm.MessageRepository.CreateMessage(ctx, message)

	if err != nil {
		return 0, errors.New("something went wrong while storing the message")
	}

	return id, nil
}
