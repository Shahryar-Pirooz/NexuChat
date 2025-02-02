package port

import (
	"context"
	"nexu-chat/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUser(ctx context.Context, id domain.UserID) (*domain.User, error)
	FilterUser(ctx context.Context, filter domain.FilterUser, limit, page uint) ([]domain.User, error)
	UpdateUser(ctx context.Context, id domain.UserID, user domain.User) error
	DeleteUser(ctx context.Context, id domain.UserID) error
}
