package domain

import (
	"errors"
	messageDomain "nexu-chat/internal/message/domain"
	userDomain "nexu-chat/internal/user/domain"

	"github.com/google/uuid"
)

var (
	ErrEmptyChatroomName   = errors.New("chatroom name cannot be empty")
	ErrChatroomNameTooLong = errors.New("chatroom name cannot exceed 100 characters")
	ErrInvalidOwner        = errors.New("invalid chatroom owner")
	ErrInvalidPageSize     = errors.New("invalid page size")
	ErrInvalidPage         = errors.New("invalid page number")
)

type ChatroomID = uuid.UUID
type Chatroom struct {
	ID       ChatroomID
	Name     string
	Messages []messageDomain.MessageID
	Users    []userDomain.UserID
	Owner    userDomain.UserID
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
