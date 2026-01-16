package service

import (
	"chat/internal/entities"
	"chat/internal/repository"
	"context"
)

type MessageService interface {
	Send(ctx context.Context, chatID int, text string) (*entities.Message, error)
}

type msgSvc struct {
	msgRepo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &msgSvc{repo}
}

func (m *msgSvc) Send(ctx context.Context, chatID int, text string) (*entities.Message, error) {
	return nil, nil
}
