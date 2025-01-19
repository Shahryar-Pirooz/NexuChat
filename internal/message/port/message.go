package port

import (
	"context"
	"nexu-chat/internal/message/domain"
)

type Repo interface {
	Create(ctx context.Context, record domain.Message) (domain.MessageID, error)
	Update(ctx context.Context, id domain.MessageID, newRecord domain.Message) error
	Read(ctx context.Context, page, limit uint, filter *domain.MessageFilter) error
	Delete(ctx context.Context, id domain.MessageID) error
}
