package http

import (
	"chat/internal/service"
	"chat/internal/transport/http/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		return
	}

	resp := dto.ChatResponse{
		ID:        chat.ID,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *ChatHandler) GetChatByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	var limit int
	switch limitStr {
	case "":
		limit = 20
	default:
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "ivalid limit parameter", http.StatusBadRequest)
			return
		}

		if limit < 0 || limit > 100 {
			limit = 20
		}
	}

	chat, err := c.chatSvc.GetByID(r.Context(), id, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.ChatResponse{
		ID:        chat.ID,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}

	for _, msg := range chat.Messages {
		msgResp := dto.MessageResponse{
			ID:        msg.ID,
			ChatID:    msg.ChatID,
			Text:      msg.Text,
			CreatedAt: msg.CreatedAt,
		}

		resp.Messages = append(resp.Messages, msgResp)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *ChatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.chatSvc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
