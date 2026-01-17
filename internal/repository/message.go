package repository

import (
	"chat/internal/db/models"
	"chat/internal/entities"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *entities.Message) (*entities.Message, error)
	GetLastByChatID(ctx context.Context, id, limit int) ([]*entities.Message, error)
}

type msgRepo struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &msgRepo{
		db: db,
	}
}

func (m *msgRepo) Create(ctx context.Context, msg *entities.Message) (*entities.Message, error) {
	model := &models.MessageModel{
		ChatID: msg.ChatID,
		Text:   msg.Text,
	}

	if err := m.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	created := &entities.Message{
		ID:        model.ID,
		ChatID:    model.ChatID,
		Text:      model.Text,
		CreatedAt: model.CreatedAt,
	}

	return created, nil
}

func (m *msgRepo) GetLastByChatID(ctx context.Context, id, limit int) ([]*entities.Message, error) {
	var models []*models.MessageModel

	if err := m.db.WithContext(ctx).Where("chat_id = ?", id).Order("created_at DESC").Limit(limit).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get last messages by id: %v", err)
	}

	msgs := make([]*entities.Message, 0, len(models))

	for _, model := range models {
		msg := &entities.Message{
			ID:        model.ID,
			ChatID:    model.ChatID,
			Text:      model.Text,
			CreatedAt: model.CreatedAt,
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}
