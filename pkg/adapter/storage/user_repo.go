package storage

import (
	"context"
	"nexu-chat/internal/user/domain"
	userPort "nexu-chat/internal/user/port"
	"nexu-chat/pkg/adapter/storage/mapper"
	"nexu-chat/pkg/adapter/storage/types"
	"nexu-chat/pkg/fp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, record domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(record)
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return uuid.Nil, err
	}
	id, err := uuid.Parse(user.ID)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil

}
func (r *userRepo) GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	var user types.User
	if err := r.db.WithContext(ctx).Where("id = ?", id.String()).First(user).Error; err != nil {
		return nil, err
	}
	domainUser := mapper.UserStorage2SDomain(user)
	return &domainUser, nil
}
func (r *userRepo) FilterUser(ctx context.Context, filter domain.FilterUser, limit, page uint) ([]domain.User, error) {
	var users []types.User
	offset := (page - 1) * limit
	query := r.db.WithContext(ctx).
		Offset(int(offset)).
		Limit(int(limit))

	if filter.Username != "" {
		query = query.Where("username = ?", filter.Username)
	}

	if filter.Role != 0 {
		query = query.Where("role = ?", filter.Role)
	}

	if err := query.Find(&users).Error; err != nil {
		return []domain.User{}, err
	}

	domainUsers := fp.Map(users, mapper.UserStorage2SDomain)
	return domainUsers, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, id domain.UserID, user domain.User) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Updates(mapper.UserDomain2Storage(user)).
		Error
}

func (r *userRepo) DeleteUser(ctx context.Context, id domain.UserID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Delete(&types.User{}).
		Error
}
