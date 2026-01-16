package service

import "chat/internal/repository"

type Service struct {
	ChatService
	MessageService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ChatService:    NewChatService(repo.ChatRepository),
		MessageService: NewMessageService(repo.MessageRepository),
	}
}
