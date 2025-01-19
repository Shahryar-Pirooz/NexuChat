package chatroom

import (
	"context"
	"fmt"
	"nexu-chat/internal/chatroom/domain"
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
	id, err := s.repo.CreateChatroom(ctx, record)
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
	err := s.repo.UpdateChatroom(ctx, id, newRecord)
	return err
}
func (s *service) GetAllChatrooms(ctx context.Context) ([]domain.Chatroom, error) {
	chatrooms, err := s.repo.FilterChatroom(ctx, &chatroomDomain.ChatroomFilter{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all chatrooms: %w", err)
	}
	return chatrooms, nil
}
func (s *service) FilterChatrooms(ctx context.Context, page, limit uint, filter *domain.ChatroomFilter) ([]domain.Chatroom, error) {
	chatrooms, err := s.repo.FilterChatroom(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get all chatrooms: %w", err)
	}
	return chatrooms, nil
}
func (s *service) DeleteChatroom(ctx context.Context, id domain.ChatroomID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	err := s.repo.DeleteChatroom(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete chatroom: %w", err)
	}
	return nil
}

// message methods
func (s *service) CreateMessage(ctx context.Context, chatroomID domain.ChatroomID, message domain.Message) (domain.MessageID, error) {
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	if err := message.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate message: %w", err)
	}
	id, err := s.repo.CreateMessage(ctx, chatroomID, message)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create message: %w", err)
	}
	return id, nil
}
func (s *service) DeleteMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID) error {
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	if err := uuid.Validate(messageID.String()); err != nil {
		return fmt.Errorf("failed to validate message id: %w", err)
	}
	err := s.repo.DeleteMessage(ctx, chatroomID, messageID)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}
func (s *service) GetAllMessages(ctx context.Context, chatroomID domain.ChatroomID) ([]domain.Message, error) {
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return nil, fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	messages, err := s.repo.GetAllMessages(ctx, chatroomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all messages: %w", err)
	}
	return messages, nil
}
func (s *service) UpdateMessage(ctx context.Context, chatroomID domain.ChatroomID, messageID domain.MessageID, newMessage domain.Message) error {
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return fmt.Errorf("failed to validate chatroom id: %w", err)
	}
	if err := uuid.Validate(messageID.String()); err != nil {
		return fmt.Errorf("failed to validate message id: %w", err)
	}
	if err := newMessage.Validate(); err != nil {
		return fmt.Errorf("failed to validate message: %w", err)
	}
	err := s.repo.UpdateMessage(ctx, chatroomID, messageID, newMessage)
	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}
	return nil
}
