package application

import (
	"context"

	"github.com/challenge/pkg/domain"
	"github.com/challenge/pkg/helpers"
)

const GetMessagesQueryName = "GetMessagesQuery"

type GetMessagesQuery struct {
	Recipient int64
	Start     int64
}

type GetMessagesQueryHandler struct {
	MessageRepository MessageRepository
}

func (msgQuery GetMessagesQuery) GetQueryName() string {
	return GetMessagesQueryName
}

func CreateGetMessagesQueryHandler(messageRepository MessageRepository) GetMessagesQueryHandler {
	return GetMessagesQueryHandler{
		MessageRepository: messageRepository,
	}
}

func (qry GetMessagesQueryHandler) Handle(ctx context.Context, qy Query) ([]domain.Message, error) {
	q, ok := qy.(GetMessagesQuery)

	if !ok {
		return nil, helpers.NewInternalServerError("Wrong command")
	}

	messages, err := qry.MessageRepository.GetMessages(ctx, q.Recipient, q.Start)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
