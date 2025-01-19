package chatroom

import (
	"context"
	"fmt"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
	chatroomPort "nexu-chat/internal/chatroom/port"

	"github.com/google/uuid"
)

type service struct {
	repo chatroomPort.Repo
}

func NewService(repo chatroomPort.Repo) chatroomPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateChatroom(ctx context.Context, record chatroomDomain.Chatroom) (chatroomDomain.ChatroomID, error) {
	if err := record.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate chatroom: %w", err)
	}
	id, err := s.repo.Create(ctx, record)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create chatroom: %w", err)
	}
	return id, nil
}
func (s *service) UpdateChatroom(ctx context.Context, id chatroomDomain.ChatroomID, newRecord chatroomDomain.Chatroom) error {
	if err := newRecord.Validate(); err != nil {
		return fmt.Errorf("failed to validate chatroom: %w", err)
	}
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	err := s.repo.Update(ctx, id, newRecord)
	return err
}
func (s *service) GetAllChatrooms(ctx context.Context) ([]chatroomDomain.Chatroom, error) {
	panic("v any")
}
func (s *service) FilterChatroom(ctx context.Context, page, limit uint, filter *chatroomDomain.ChatroomFilter) ([]chatroomDomain.Chatroom, error) {
	panic("v any")
}
func (s *service) Delete(ctx context.Context, id chatroomDomain.ChatroomID) error {
	panic("v any")
}
