package chatroom

import (
	"context"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
	chatroomPort "nexu-chat/internal/chatroom/port"
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
	panic("v any")
}
func (s *service) UpdateChatroom(ctx context.Context, id chatroomDomain.ChatroomID, newRecord chatroomDomain.Chatroom) error {
	panic("v any")
}
func (s *service) GetChatroomByID(ctx context.Context, id chatroomDomain.ChatroomID) (*chatroomDomain.Chatroom, error) {
	panic("v any")
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
