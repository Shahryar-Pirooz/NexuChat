package storage

import (
	"context"
	"fmt"
	"nexu-chat/internal/user/domain"
	userDomain "nexu-chat/internal/user/domain"
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
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(ctx context.Context, record userDomain.User) (userDomain.UserID, error) {
	if err := record.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("validation error: %w", err)
	}
	chatroom := mapper.UserDomain2Storage(record)
	result := ur.db.WithContext(ctx).Create(&chatroom)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", result.Error)
	}
	id, err := uuid.Parse(chatroom.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", result.Error)
	}
	return id, nil
}
func (ur *userRepo) UpdateUser(ctx context.Context, userID userDomain.UserID, newRecord userDomain.User) error {
	user := new(types.User)
	if err := uuid.Validate(userID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	if err := newRecord.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	user = mapper.UserDomain2Storage(newRecord)
	result := ur.db.WithContext(ctx).Where("id = ?", userID.String()).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}

	return nil
}
func (ur *userRepo) GetUserByID(ctx context.Context, userID userDomain.UserID) (*domain.User, error) {
	user := new(types.User)
	if err := uuid.Validate(userID.String()); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	result := ur.db.WithContext(ctx).Where("id = ?", userID.String()).Find(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}
	userDomain := mapper.UserStorage2Domain(*user)
	return userDomain, nil
}
func (ur *userRepo) FilterUser(ctx context.Context, filter *userDomain.FilterUser) ([]userDomain.User, error) {
	var users []types.User
	query := ur.db
	if filter != nil {
		if filter.Username != "" {
			query = query.Where("username LIKE ?", "%"+filter.Username+"%").Find(&users)
		}
		if filter.IP != "" {
			query = query.Where("ip LIKE ?", "%"+filter.IP+"%").Find(&users)
		}
		if filter.Connected {
			query = query.Where("connect = ?", "true").Find(&users)
		}
	}
	usersDomain := fp.Map(users, mapper.UserStorage2Domain)

	return usersDomain, nil
}
func (ur *userRepo) DeleteUser(ctx context.Context, userID userDomain.UserID) error {
	user := new(types.User)
	if err := uuid.Validate(userID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	result := ur.db.WithContext(ctx).Where("id = ?", userID.String()).Delete(user)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	return nil
}
