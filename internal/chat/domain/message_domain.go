package domain

import (
	userDomain "nexu-chat/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type MessageID = uuid.UUID
type Message struct {
	ID         MessageID
	ChatroomID ChatroomID
	UserID     userDomain.UserID
	Content    string
	CreatedAt  time.Time
	UpdateAt   time.Time
}

type FilterMessage struct {
	ChatroomID ChatroomID
}

func (m *Message) Validate() error {
	if m.Content == "" {
		return ErrEmptyContent
	}

	if m.ChatroomID == uuid.Nil {
		return ErrInvalidChatroomID
	}

	if m.UserID == uuid.Nil {
		return ErrInvalidUserID
	}

	if m.CreatedAt.IsZero() {
		return ErrInvalidCreation
	}

	if m.UpdateAt.IsZero() {
		return ErrInvalidUpdate
	}

	return nil
}
