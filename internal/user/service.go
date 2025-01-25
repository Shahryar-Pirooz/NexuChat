package user

import (
	"context"
	"fmt"
	userDomain "nexu-chat/internal/user/domain"
	userPort "nexu-chat/internal/user/port"

	"github.com/google/uuid"
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
	if err := record.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate user: %w", err)
	}
	id, err := s.repo.CreateUser(ctx, record)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}
func (s *service) UpdateUser(ctx context.Context, id userDomain.UserID, newRecord userDomain.User) error {
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("failed to validate user id: %w", err)
	}
	if err := newRecord.Validate(); err != nil {
		return fmt.Errorf("failed to validate user: %w", err)
	}
	err := s.repo.UpdateUser(ctx, id, newRecord)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
func (s *service) GetAllUsers(ctx context.Context) ([]userDomain.User, error) {
	users, err := s.repo.FilterUser(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}
func (s *service) GetAllActiveUser(ctx context.Context) ([]userDomain.User, error) {
	filter := &userDomain.FilterUser{
		Connected: true,
	}
	users, err := s.repo.FilterUser(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get all active users: %w", err)
	}
	return users, nil
}
func (s *service) AuthenticateUser(ctx context.Context, username, password string) (*userDomain.User, error) {
	filter := &userDomain.FilterUser{
		Username: username,
	}
	users, err := s.repo.FilterUser(ctx, filter)
	user := &users[0]
	if err != nil {
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	if !user.ComparePassword(password) {
		return nil, nil
	}
	return user, nil
}

func (s *service) DeleteUser(ctx context.Context, id userDomain.UserID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("failed to validate user id: %w", err)
	}
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
