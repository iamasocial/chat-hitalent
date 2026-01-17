package entities

import "time"

type Chat struct {
	ID        int
	Title     string
	CreatedAt time.Time
	Messages  []*Message
}
