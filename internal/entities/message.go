package entities

import "time"

type Message struct {
	ID        int64
	ChatID    int64
	Text      string
	CreatedAt time.Time
}
