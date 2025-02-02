package domain

import (
	"time"

	"github.com/google/uuid"
)

type ChatroomID = uuid.UUID

type Chatroom struct {
	ID        ChatroomID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FilterChatroom struct {
	Name string
}

func (c *Chatroom) Validate() error {
	if err := uuid.Validate(c.ID.String()); err != nil {
		return ErrInvalidChatroomID
	}
	if c.Name == "" {
		return ErrEmptyName
	}

	if len(c.Name) > 50 {
		return ErrNameTooLong
	}

	if c.CreatedAt.IsZero() {
		return ErrInvalidCreation
	}

	if c.UpdatedAt.IsZero() {
		return ErrInvalidUpdate
	}

	return nil
}
