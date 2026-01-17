package repository

import (
	"chat/internal/db/models"
	"chat/internal/entities"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *entities.Chat) (*entities.Chat, error)
	GetByID(ctx context.Context, id int) (*entities.Chat, error)
	Delete(ctx context.Context, id int) error
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

func (c *chatRepo) GetByID(ctx context.Context, id int) (*entities.Chat, error) {
	model := &models.ChatModel{
		ID: id,
	}

	if err := c.db.WithContext(ctx).First(model, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get chat by ID: %v", err)
	}

	chat := &entities.Chat{
		ID:        model.ID,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
	}

	return chat, nil
}

func (c *chatRepo) Delete(ctx context.Context, id int) error {
	if err := c.db.Delete(&models.ChatModel{}, id); err != nil {
		return err.Error
	}

	return nil
}
