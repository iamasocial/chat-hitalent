package http

import (
	"chat/internal/service"
	"chat/internal/transport/http/dto"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageHandler struct {
	msgSvc service.MessageService
}

func NewMessageHandler(svc service.MessageService) *MessageHandler {
	return &MessageHandler{msgSvc: svc}
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	var req dto.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	msg, err := m.msgSvc.Send(r.Context(), id, req.Text)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmptyMessage),
			errors.Is(err, service.ErrMessageTooLong):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, service.ErrChatNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	resp := dto.MessageResponse{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
