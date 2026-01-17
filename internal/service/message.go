package service

import (
	"chat/internal/entities"
	"chat/internal/repository"
	"context"
	"errors"
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
		return nil, ErrEmptyMessage
	}

	if len(text) > 5000 {
		return nil, ErrMessageTooLong
	}

	msg := &entities.Message{
		ChatID: chatID,
		Text:   text,
	}

	created, err := m.msgRepo.Create(ctx, msg)
	if err != nil {
		if errors.Is(err, ErrChatNotFound) {
			return nil, ErrChatNotFound
		}
	}

	return created, nil
}
