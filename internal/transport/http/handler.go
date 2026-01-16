package http

import "chat/internal/service"

type Handler struct {
	chat    *ChatHandler
	message *MessageHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		chat:    NewChatHandler(svc.ChatService),
		message: NewMessageHandler(svc.MessageService),
	}
}
