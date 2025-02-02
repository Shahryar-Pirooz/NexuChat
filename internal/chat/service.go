package chat

import (
	"context"
	"nexu-chat/internal/chat/domain"
	"nexu-chat/internal/chat/port"

	"github.com/google/uuid"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateChatroom(ctx context.Context, chatroom domain.Chatroom) (domain.ChatroomID, error) {
	if err := chatroom.Validate(); err != nil {
		return uuid.Nil, err
	}
	return s.repo.CreateChatroom(ctx, chatroom)
}
func (s *service) GetChatroom(ctx context.Context, id domain.ChatroomID) (*domain.Chatroom, error) {
	if err := uuid.Validate(id.String()); err != nil {
		return nil, err
	}
	return s.repo.GetChatroomByID(ctx, id)
}
func (s *service) FilterChatroom(ctx context.Context, filter domain.FilterChatroom, limit, page uint) ([]domain.Chatroom, error) {
	return s.repo.FilterChatroom(ctx, filter, limit, page)

}
func (s *service) UpdateChatroom(ctx context.Context, id domain.ChatroomID, chatroom domain.Chatroom) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	if err := chatroom.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateChatroom(ctx, id, chatroom)
}
func (s *service) DeleteChatroom(ctx context.Context, id domain.ChatroomID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	return s.repo.DeleteChatroom(ctx, id)
}
func (s *service) CreateMessage(ctx context.Context, message domain.Message) (domain.MessageID, error) {
	if err := message.Validate(); err != nil {
		return uuid.Nil, err
	}
	return s.repo.CreateMessage(ctx, message)
}
func (s *service) GetMessage(ctx context.Context, id domain.MessageID) (*domain.Message, error) {
	if err := uuid.Validate(id.String()); err != nil {
		return nil, err
	}
	return s.repo.GetMessageByID(ctx, id)
}
func (s *service) FilterMessage(ctx context.Context, filter domain.FilterMessage, limit, page uint) ([]domain.Message, error) {
	return s.repo.FilterMessage(ctx, filter, limit, page)
}
func (s *service) UpdateMessage(ctx context.Context, id domain.MessageID, message domain.Message) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	if err := message.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateMessage(ctx, id, message)
}
func (s *service) DeleteMessage(ctx context.Context, id domain.MessageID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	return s.repo.DeleteMessage(ctx, id)
}
