package service

import "errors"

var (
	ErrEmptyTitle     = errors.New("chat title is empty")
	ErrTitleTooLong   = errors.New("chat title exceeds 200 characters")
	ErrChatNotFound   = errors.New("chat not found")
	ErrEmptyMessage   = errors.New("message is empty")
	ErrMessageTooLong = errors.New("message exceeds 5000 characters")
)
