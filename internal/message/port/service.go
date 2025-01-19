package port

import (
	"context"
	"nexu-chat/internal/message/domain"
)

type Service interface {
	CreateMessage(ctx context.Context, record domain.Message) (domain.MessageID, error)
	UpdateMessage(ctx context.Context, id domain.MessageID, newRecord domain.Message) error
	GetAllMessages(ctx context.Context) ([]domain.Message, error)
	DeleteMessage(ctx context.Context, id domain.MessageID) error
}
