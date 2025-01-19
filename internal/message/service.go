package message

import (
	"context"
	messageDomain "nexu-chat/internal/message/domain"
	messagePort "nexu-chat/internal/message/port"
)

type service struct {
	repo messagePort.Repo
}

func NewService(repo messagePort.Repo) messagePort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateMessage(ctx context.Context, record messageDomain.Message) (messageDomain.MessageID, error) {
	panic("v any")
}
func (s *service) UpdateMessage(ctx context.Context, id messageDomain.MessageID, newRecord messageDomain.Message) error {
	panic("v any")
}
func (s *service) GetAllMessages(ctx context.Context) ([]messageDomain.Message, error) {
	panic("v any")
}
func (s *service) DeleteMessage(ctx context.Context, id messageDomain.MessageID) error {
	panic("v any")
}
