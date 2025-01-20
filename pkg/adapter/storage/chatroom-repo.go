package storage

import (
	"context"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
	chatroomPort "nexu-chat/internal/chatroom/port"

	"gorm.io/gorm"
)

type chatroomRepo struct {
	db *gorm.DB
}

func NewChatroomRepo(db *gorm.DB) chatroomPort.Repo {
	return &chatroomRepo{db: db}
}

func (cp *chatroomRepo) CreateChatroom(ctx context.Context, record chatroomDomain.Chatroom) (chatroomDomain.ChatroomID, error) {
	panic("v any")
}
func (cp *chatroomRepo) DeleteChatroom(ctx context.Context, chatroomID chatroomDomain.ChatroomID) error {
	panic("v any")
}
func (cp *chatroomRepo) FilterChatroom(ctx context.Context, filter *chatroomDomain.ChatroomFilter) ([]chatroomDomain.Chatroom, error) {
	panic("v any")
}
func (cp *chatroomRepo) GetChatroomByID(ctx context.Context, chatroomID chatroomDomain.ChatroomID) (*chatroomDomain.Chatroom, error) {
	panic("v any")
}
func (cp *chatroomRepo) UpdateChatroom(ctx context.Context, chatroomID chatroomDomain.ChatroomID, newRecord chatroomDomain.Chatroom) error {
	panic("v any")
}

// message methods
func (cp *chatroomRepo) CreateMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, message chatroomDomain.Message) (chatroomDomain.MessageID, error) {
	panic("v any")
}
func (cp *chatroomRepo) DeleteMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID) error {
	panic("v any")
}
func (cp *chatroomRepo) GetAllMessages(ctx context.Context, chatroomID chatroomDomain.ChatroomID) ([]chatroomDomain.Message, error) {
	panic("v any")
}
func (cp *chatroomRepo) GetMessageByID(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID) (*chatroomDomain.Message, error) {
	panic("v any")
}
func (cp *chatroomRepo) UpdateMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID, newMessage chatroomDomain.Message) error {
	panic("v any")
}
