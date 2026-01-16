package http

import (
	"chat/internal/service"
	"chat/internal/transport/http/dto"
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	chatSvc service.ChatService
}

func NewChatHandler(svc service.ChatService) *ChatHandler {
	return &ChatHandler{chatSvc: svc}
}

func (c *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err := c.chatSvc.CreateChat(r.Context(), req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := dto.ChatResponse{
		ID:        chat.ID,
		Title:     chat.Title,
		CreatedAd: chat.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
