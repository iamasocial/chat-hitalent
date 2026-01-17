package models

import "time"

type MessageModel struct {
	ID        int       `gorm:"primaryKey"`
	ChatID    int       `gorm:"not null;index"`
	Text      string    `gorm:"type:text;not null;check:char_length(text) <= 5000"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (MessageModel) TableName() string {
	return "messages"
}
