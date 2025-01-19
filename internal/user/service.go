package user

import (
	"context"
	userDomain "nexu-chat/internal/user/domain"
	userPort "nexu-chat/internal/user/port"
)

type service struct {
	repo userPort.Repo
}

func NewService(repo userPort.Repo) userPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, record userDomain.User) (userDomain.UserID, error) {
	panic("v any")
}
func (s *service) UpdateUser(ctx context.Context, id userDomain.UserID, newRecord userDomain.User) error {
	panic("v any")
}
func (s *service) GetAllUsers(ctx context.Context) ([]userDomain.User, error) {
	panic("v any")
}
func (s *service) GetAllActiveUser(ctx context.Context) ([]userDomain.User, error) {
	panic("v any")
}
func (s *service) DeleteUser(ctx context.Context, id userDomain.UserID) error {
	panic("v any")
}
