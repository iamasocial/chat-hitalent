package entities

import "time"

type Message struct {
	ID        int
	ChatID    int
	Text      string
	CreatedAt time.Time
}
