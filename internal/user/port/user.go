package port

import (
	"context"
	"nexu-chat/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, record domain.User) (domain.UserID, error)
	Update(ctx context.Context, id domain.UserID, newRecord domain.User) error
	Read(ctx context.Context, page, limit uint, id domain.UserID) error
	Delete(ctx context.Context, id domain.UserID) error
}
