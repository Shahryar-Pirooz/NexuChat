package domain

import (
	userDomain "nexu-chat/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       MessageID
	SenderID userDomain.UserID
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}

type MessageFilter struct {
	SenderID userDomain.UserID
	FromDate time.Time
	ToDate   time.Time
	Content  string
}


func (m *Message) Validate() error {
	if m.Content == "" {
		return ErrEmptyContent
	}

	if m.SenderID == uuid.Nil {
		return ErrInvalidSenderID
	}

	return nil
}
