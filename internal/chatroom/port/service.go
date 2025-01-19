package port

import (
	"context"
	"nexu-chat/internal/chatroom/domain"
)

type Service interface {
	CreateChatroom(ctx context.Context, record domain.Chatroom) (domain.ChatroomID, error)
	UpdateChatroom(ctx context.Context, id domain.ChatroomID, newRecord domain.Chatroom) error
	GetAllChatrooms(ctx context.Context) ([]domain.Chatroom, error)
	FilterChatrooms(ctx context.Context, page, limit uint, filter *domain.ChatroomFilter) ([]domain.Chatroom, error)
	DeleteChatroom(ctx context.Context, id domain.ChatroomID) error
	// message methods
	CreateMessage(ctx context.Context, chatroomID domain.ChatroomID, message domain.Message) (domain.MessageID, error)
	DeleteMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID) error
	GetAllMessages(ctx context.Context, chatroomID domain.ChatroomID) ([]domain.Message, error)
	UpdateMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID, newMessage domain.Message) error
}
