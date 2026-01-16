package http

import "chat/internal/service"

type MessageHandler struct {
	msgSvc service.MessageService
}

func NewMessageHandler(svc service.MessageService) *MessageHandler {
	return &MessageHandler{msgSvc: svc}
}
