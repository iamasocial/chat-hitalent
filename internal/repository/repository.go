package repository

import "gorm.io/gorm"

type Repository struct {
	ChatRepository
	MessageRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ChatRepository:    NewChatRepository(db),
		MessageRepository: NewMessageRepository(db),
	}
}
