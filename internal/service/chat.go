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
}

type chatSvc struct {
	chatRepo repository.ChatRepository
}

func NewChatService(chatRepo repository.ChatRepository) ChatService {
	return &chatSvc{
		chatRepo: chatRepo,
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
