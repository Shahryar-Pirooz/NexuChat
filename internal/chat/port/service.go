package port

import (
	"context"
	"nexu-chat/internal/chat/domain"
)

type Service interface {
	CreateChatroom(ctx context.Context, chatroom domain.Chatroom) (domain.ChatroomID, error)
	GetChatroom(ctx context.Context, id domain.ChatroomID) (*domain.Chatroom, error)
	FilterChatroom(ctx context.Context, filter domain.FilterChatroom, limit, page uint) ([]domain.Chatroom, error)
	UpdateChatroom(ctx context.Context, id domain.ChatroomID, chatroom domain.Chatroom) error
	DeleteChatroom(ctx context.Context, id domain.ChatroomID) error
	CreateMessage(ctx context.Context, message domain.Message) (domain.MessageID, error)
	GetMessage(ctx context.Context, id domain.MessageID) (*domain.Message, error)
	FilterMessage(ctx context.Context, filter domain.FilterMessage, limit, page uint) ([]domain.Message, error)
	UpdateMessage(ctx context.Context, id domain.MessageID, message domain.Message) error
	DeleteMessage(ctx context.Context, id domain.MessageID) error
}
