package models

import "time"

type ChatModel struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (ChatModel) TableName() string {
	return "chats"
}
