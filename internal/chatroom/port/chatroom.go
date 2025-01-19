package port

import (
	"context"
	"nexu-chat/internal/chatroom/domain"
)

type Repo interface {
	CreateChatroom(ctx context.Context, record domain.Chatroom) (domain.ChatroomID, error)
	DeleteChatroom(ctx context.Context, chatroomID domain.ChatroomID) error
	FilterChatroom(ctx context.Context, filter *domain.ChatroomFilter) ([]domain.Chatroom, error)
	GetChatroomByID(ctx context.Context, chatroomID domain.ChatroomID) (*domain.Chatroom, error)
	UpdateChatroom(ctx context.Context, chatroomID domain.ChatroomID, newRecord domain.Chatroom) error
	// message methods
	CreateMessage(ctx context.Context, chatroomID domain.ChatroomID, message domain.Message) (domain.MessageID, error)
	DeleteMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID) error
	GetAllMessages(ctx context.Context, chatroomID domain.ChatroomID) ([]domain.Message, error)
	GetMessageByID(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID) (*domain.Message, error)
	UpdateMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID, newMessage domain.Message) error
}
