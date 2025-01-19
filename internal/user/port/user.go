package port

import (
	"context"
	"nexu-chat/internal/user/domain"
)

type Repo interface {
	CreateUser(ctx context.Context, record domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, userID domain.UserID, newRecord domain.User) error
	ReadUser(ctx context.Context, userID domain.UserID) error
	FilterUser(ctx context.Context, filter *domain.FilterUser) ([]domain.User, error)
	DeleteUser(ctx context.Context, userID domain.UserID) error
}
