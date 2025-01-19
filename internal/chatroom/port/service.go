package port

import (
	"context"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
)

type Service interface {
	CreateChatroom(ctx context.Context, record chatroomDomain.Chatroom) (chatroomDomain.ChatroomID, error)
	UpdateChatroom(ctx context.Context, id chatroomDomain.ChatroomID, newRecord chatroomDomain.Chatroom) error
	GetAllChatrooms(ctx context.Context) ([]chatroomDomain.Chatroom, error)
	FilterChatroom(ctx context.Context, page, limit uint, filter *chatroomDomain.ChatroomFilter) ([]chatroomDomain.Chatroom, error)
	Delete(ctx context.Context, id chatroomDomain.ChatroomID) error
}
