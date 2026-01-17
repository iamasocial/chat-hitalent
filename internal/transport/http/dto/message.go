package dto

import "time"

type SendMessageRequest struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
