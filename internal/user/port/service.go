package port

import (
	"context"
	"nexu-chat/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, record domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, id domain.UserID, newRecord domain.User) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetAllActiveUser(ctx context.Context) ([]domain.User, error)
	AuthenticateUser(ctx context.Context, username, password string) (*domain.User, error)
	DeleteUser(ctx context.Context, id domain.UserID) error
}
