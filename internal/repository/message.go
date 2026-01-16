package repository

import (
	"chat/internal/entities"
	"context"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *entities.Message) (*entities.Message, error)
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
	return nil, nil
}
