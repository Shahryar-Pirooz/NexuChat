package storage

import (
	"context"
	userDomain "nexu-chat/internal/user/domain"
	userPort "nexu-chat/internal/user/port"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(ctx context.Context, record userDomain.User) (userDomain.UserID, error) {
	panic("v any")
}
func (ur *userRepo) UpdateUser(ctx context.Context, userID userDomain.UserID, newRecord userDomain.User) error {
	panic("v any")
}
func (ur *userRepo) ReadUser(ctx context.Context, userID userDomain.UserID) error {
	panic("v any")
}
func (ur *userRepo) FilterUser(ctx context.Context, filter *userDomain.FilterUser) ([]userDomain.User, error) {
	panic("v any")
}
func (ur *userRepo) DeleteUser(ctx context.Context, userID userDomain.UserID) error {
	panic("v any")
}
