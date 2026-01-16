package models

import "time"

type MessageModel struct {
	ID        int64     `gorm:"primaryKey"`
	ChatID    int64     `gorm:"not null;index"`
	Text      string    `gorm:"type:text;not null;check:char_length(text) <= 5000"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (MessageModel) TableName() string {
	return "messages"
}
