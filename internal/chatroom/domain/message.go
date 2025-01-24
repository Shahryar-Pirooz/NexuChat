package domain

import (
	"nexu-chat/internal/user/domain"
	userDomain "nexu-chat/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         MessageID
	SenderID   userDomain.UserID
	ChatroomID ChatroomID
	Content    string
	CreateAt   time.Time
	UpdateAt   time.Time
}

type MessageFilter struct {
	SenderID   userDomain.UserID
	ChatroomID ChatroomID
	FromDate   time.Time
	ToDate     time.Time
	Content    string
}

type MessageBroker interface {
	PublishMessage(msg Message) error
	SubscribeMessages(handler func(Message)) error
	PublishUserJoined(user domain.User) error
	SubscribeUserJoined(handler func(domain.User)) error
	PublishUserLeft(user domain.User) error
	SubscribeUserLeft(handler func(domain.User)) error
	Close()
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
