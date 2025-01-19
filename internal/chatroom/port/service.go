package port

import (
	"context"
	"nexu-chat/internal/chatroom/domain"
)

type Service interface {
	CreateChatroom(ctx context.Context, record domain.Chatroom) (domain.ChatroomID, error)
	UpdateChatroom(ctx context.Context, id domain.ChatroomID, newRecord domain.Chatroom) error
	GetChatroomByID(ctx context.Context, id domain.ChatroomID) (*domain.Chatroom, error)
	GetAllChatrooms(ctx context.Context) ([]domain.Chatroom, error)
	FilterChatroom(ctx context.Context, page, limit uint, filter *domain.ChatroomFilter) ([]domain.Chatroom, error)
	Delete(ctx context.Context, id domain.ChatroomID) error
}
