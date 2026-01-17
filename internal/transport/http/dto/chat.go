package dto

import "time"

type CreateChatRequest struct {
	Title string `json:"title"`
}

type ChatResponse struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	CreatedAt time.Time         `json:"created_at"`
	Messages  []MessageResponse `json:"messages,omitempty"`
}
