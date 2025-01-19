package domain

import (
	"errors"
	userDomain "nexu-chat/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyChatroomName   = errors.New("chatroom name cannot be empty")
	ErrChatroomNameTooLong = errors.New("chatroom name cannot exceed 100 characters")
	ErrInvalidOwner        = errors.New("invalid chatroom owner")
	ErrInvalidPageSize     = errors.New("invalid page size")
	ErrInvalidPage         = errors.New("invalid page number")
	ErrEmptyContent        = errors.New("message content cannot be empty")
	ErrInvalidSenderID     = errors.New("invalid sender ID")
)

type (
	ChatroomID = uuid.UUID
	MessageID  = uuid.UUID
)

type Chatroom struct {
	ID       ChatroomID
	Name     string
	Messages []MessageID
	Users    []userDomain.UserID
	Owner    userDomain.UserID
	CreateAt time.Time
	UpdateAt time.Time
}

type ChatroomFilter struct {
	Name     string
	OwnerID  userDomain.UserID
	Page     uint
	PageSize uint
}

func (c *Chatroom) Validate() error {
	if c.Name == "" {
		return ErrEmptyChatroomName
	}
	if len(c.Name) > 100 {
		return ErrChatroomNameTooLong
	}
	if c.Owner == uuid.Nil {
		return ErrInvalidOwner
	}
	return nil
}
