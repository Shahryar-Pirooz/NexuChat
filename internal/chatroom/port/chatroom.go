package port

import (
	"context"
	"nexu-chat/internal/chatroom/domain"
)

type Repo interface {
	Create(ctx context.Context, record domain.Chatroom) (domain.ChatroomID, error)
	Update(ctx context.Context, id domain.ChatroomID, newRecord domain.Chatroom) error
	Read(ctx context.Context, page, limit uint, filter *domain.ChatroomFilter) error
	Delete(ctx context.Context, id domain.ChatroomID) error
}
