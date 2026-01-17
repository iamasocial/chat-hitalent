package http

import (
	"chat/internal/service"
	"chat/internal/transport/http/dto"
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req dto.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := m.msgSvc.Send(r.Context(), id, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
