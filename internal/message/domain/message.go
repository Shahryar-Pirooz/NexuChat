package domain

import (
	"errors"
	userDomain "nexu-chat/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyContent    = errors.New("message content cannot be empty")
	ErrInvalidSenderID = errors.New("invalid sender ID")
)

type MessageID = uuid.UUID

type MessageFilter struct {
	SenderID userDomain.UserID
	FromDate time.Time
	ToDate   time.Time
	Content  string
}

type Message struct {
	ID       MessageID
	SenderID userDomain.UserID
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
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
