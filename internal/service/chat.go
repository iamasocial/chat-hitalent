package service

import (
	"chat/internal/entities"
	"chat/internal/repository"
	"context"
	"fmt"
	"strings"
)

type ChatService interface {
	CreateChat(ctx context.Context, title string) (*entities.Chat, error)
	GetByID(ctx context.Context, id, limit int) (*entities.Chat, error)
	Delete(ctx context.Context, id int) error
}

type chatSvc struct {
	chatRepo repository.ChatRepository
	msgRepo  repository.MessageRepository
}

func NewChatService(chatRepo repository.ChatRepository, msgRepo repository.MessageRepository) ChatService {
	return &chatSvc{
		chatRepo: chatRepo,
		msgRepo:  msgRepo,
	}
}

func (c *chatSvc) CreateChat(ctx context.Context, title string) (*entities.Chat, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return nil, fmt.Errorf("chat title is empty")
	}

	if len(title) > 200 {
		return nil, fmt.Errorf("chat title exceeds 200 characters")
	}

	chat := &entities.Chat{
		Title: title,
	}

	createdChat, err := c.chatRepo.Create(ctx, chat)
	if err != nil {
		return nil, err
	}

	return createdChat, nil
}

func (c *chatSvc) GetByID(ctx context.Context, id, limit int) (*entities.Chat, error) {
	chat, err := c.chatRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	msgs, err := c.msgRepo.GetLastByChatID(ctx, id, limit)
	if err != nil {
		return nil, err
	}

	chat.Messages = msgs

	return chat, nil
}

func (c *chatSvc) Delete(ctx context.Context, id int) error {
	return c.chatRepo.Delete(ctx, id)
}
