package application

import (
	"context"
	"fmt"

	"github.com/challenge/pkg/domain"
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
		fmt.Println("wrong query")
	}

	return qry.MessageRepository.GetMessages(ctx, q.Recipient, q.Start)
}
