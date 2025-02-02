package user

import (
	"context"
	"nexu-chat/internal/user/domain"
	"nexu-chat/internal/user/port"

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

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	if err := user.Validate(); err != nil {
		return uuid.Nil, err
	}
	return s.repo.CreateUser(ctx, user)
}
func (s *service) GetUser(ctx context.Context, id domain.UserID) (*domain.User, error) {
	if err := uuid.Validate(id.String()); err != nil {
		return nil, err
	}
	return s.repo.GetUserByID(ctx, id)
}
func (s *service) FilterUser(ctx context.Context, filter domain.FilterUser, limit, page uint) ([]domain.User, error) {
	return s.repo.FilterUser(ctx, filter, limit, page)
}
func (s *service) UpdateUser(ctx context.Context, id domain.UserID, user domain.User) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	if err := user.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateUser(ctx, id, user)
}
func (s *service) DeleteUser(ctx context.Context, id domain.UserID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return err
	}
	return s.repo.DeleteUser(ctx, id)
}
