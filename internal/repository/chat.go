package repository

import (
	"chat/internal/db/models"
	"chat/internal/entities"
	"context"

	"gorm.io/gorm"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *entities.Chat) (*entities.Chat, error)
}

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepo{db: db}
}

func (c *chatRepo) Create(ctx context.Context, chat *entities.Chat) (*entities.Chat, error) {
	model := &models.ChatModel{
		Title: chat.Title,
	}

	if err := c.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	created := &entities.Chat{
		ID:        model.ID,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
	}

	return created, nil
}
