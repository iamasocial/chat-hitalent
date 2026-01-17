package service

import (
	"chat/internal/entities"
	"chat/internal/repository"
	"context"
	"fmt"
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
	if text == "" {
		return nil, fmt.Errorf("message cannot be empty")
	}

	if len(text) > 5000 {
		return nil, fmt.Errorf("message exceeds 5000 characters")
	}

	msg := &entities.Message{
		ChatID: chatID,
		Text:   text,
	}

	created, err := m.msgRepo.Create(ctx, msg)
	if err != nil {
		return nil, err
	}

	return created, nil
}
