package storage

import (
	"context"
	"nexu-chat/internal/chat/domain"
	chatPort "nexu-chat/internal/chat/port"
	"nexu-chat/pkg/adapter/storage/mapper"
	"nexu-chat/pkg/adapter/storage/types"
	"nexu-chat/pkg/fp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) chatPort.Repo {
	return &chatRepo{
		db: db,
	}
}

func (r *chatRepo) CreateChatroom(ctx context.Context, record domain.Chatroom) (domain.ChatroomID, error) {
	if err := record.Validate(); err != nil {
		return uuid.Nil, err
	}

	chatroomStorage := mapper.ChatroomDomain2Storage(record)
	if err := r.db.WithContext(ctx).Create(chatroomStorage).Error; err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(chatroomStorage.ID)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *chatRepo) GetChatroomByID(ctx context.Context, id domain.ChatroomID) (*domain.Chatroom, error) {
	if err := uuid.Validate(id.String()); err != nil {
		return nil, err
	}

	var chatroom types.Chatroom
	if err := r.db.WithContext(ctx).First(&chatroom, id.String()).Error; err != nil {
		return nil, err
	}

	result := mapper.ChatroomStorage2Domain(chatroom)
	return &result, nil
}

func (r *chatRepo) FilterChatroom(ctx context.Context, filter domain.FilterChatroom, limit, page uint) ([]domain.Chatroom, error) {
	var chatrooms []types.Chatroom
	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Where("name = ?", filter.Name).
		Order("id").
		Offset(int(offset)).
		Limit(int(limit))

	if err := query.Find(&chatrooms).Error; err != nil {
		return nil, err
	}

	return fp.Map(chatrooms, mapper.ChatroomStorage2Domain), nil
}

func (r *chatRepo) UpdateChatroom(ctx context.Context, id domain.ChatroomID, newRecord domain.Chatroom) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Updates(mapper.ChatroomDomain2Storage(newRecord)).
		Error
}

func (r *chatRepo) DeleteChatroom(ctx context.Context, id domain.ChatroomID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Delete(&types.Chatroom{}).
		Error
}
func (r *chatRepo) CreateMessage(ctx context.Context, record domain.Message) (domain.MessageID, error) {
	if err := record.Validate(); err != nil {
		return uuid.Nil, err
	}

	messageStorage := mapper.MessageDomain2Storage(record)
	if err := r.db.WithContext(ctx).Create(messageStorage).Error; err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(messageStorage.ID)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *chatRepo) GetMessageByID(ctx context.Context, id domain.MessageID) (*domain.Message, error) {
	if err := uuid.Validate(id.String()); err != nil {
		return nil, err
	}

	var message types.Message
	if err := r.db.WithContext(ctx).First(&message, id.String()).Error; err != nil {
		return nil, err
	}

	result := mapper.MessageStorage2Domain(message)
	return &result, nil
}

func (r *chatRepo) FilterMessage(ctx context.Context, filter domain.FilterMessage, limit, page uint) ([]domain.Message, error) {
	var messages []types.Message
	offset := (page - 1) * limit

	query := r.db.WithContext(ctx).
		Where("chatroom_id = ?", filter.ChatroomID).
		Order("id").
		Offset(int(offset)).
		Limit(int(limit))

	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}

	return fp.Map(messages, mapper.MessageStorage2Domain), nil
}

func (r *chatRepo) UpdateMessage(ctx context.Context, id domain.MessageID, message domain.Message) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Updates(mapper.MessageDomain2Storage(message)).
		Error
}

func (r *chatRepo) DeleteMessage(ctx context.Context, id domain.MessageID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id.String()).
		Delete(&types.Message{}).
		Error
}
