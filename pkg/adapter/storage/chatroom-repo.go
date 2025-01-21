package storage

import (
	"context"
	"fmt"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
	chatroomPort "nexu-chat/internal/chatroom/port"
	"nexu-chat/pkg/adapter/storage/mapper"
	"nexu-chat/pkg/adapter/storage/types"
	"nexu-chat/pkg/fp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type chatroomRepo struct {
	db *gorm.DB
}

func NewChatroomRepo(db *gorm.DB) chatroomPort.Repo {
	return &chatroomRepo{db: db}
}

func (cp *chatroomRepo) CreateChatroom(ctx context.Context, record chatroomDomain.Chatroom) (chatroomDomain.ChatroomID, error) {
	chatroom := new(types.Chatroom)
	if err := record.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("validation error: %w", err)
	}
	chatroom = mapper.ChatroomDomain2Storage(record)
	result := cp.db.WithContext(ctx).Create(chatroom)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create chatroom: %w", result.Error)
	}
	id, err := uuid.Parse(chatroom.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create chatroom: %w", result.Error)
	}
	return id, nil
}
func (cp *chatroomRepo) DeleteChatroom(ctx context.Context, chatroomID chatroomDomain.ChatroomID) error {
	chatroom := new(types.Chatroom)
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	result := cp.db.WithContext(ctx).Where("id = ?", chatroomID.String()).Delete(chatroom)
	if result.Error != nil {
		return fmt.Errorf("failed to delte chatroom: %w", result.Error)
	}
	return nil
}
func (cp *chatroomRepo) FilterChatroom(ctx context.Context, filter *chatroomDomain.ChatroomFilter) ([]chatroomDomain.Chatroom, error) {
	var chatrooms []types.Chatroom
	query := cp.db
	if filter != nil {
		if filter.Name != "" {
			query = query.Where("name LIKE ?", "%"+filter.Name+"%").Find(&chatrooms)
		}
	}
	chatroomsDomain := fp.Map(chatrooms, mapper.ChatroomStorage2Domain)

	return chatroomsDomain, nil
}
func (cp *chatroomRepo) GetChatroomByID(ctx context.Context, chatroomID chatroomDomain.ChatroomID) (*chatroomDomain.Chatroom, error) {
	chatroom := new(types.Chatroom)
	if err := uuid.Validate(chatroomID.String()); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	result := cp.db.WithContext(ctx).Find(chatroom, chatroomID.String())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to delte chatroom: %w", result.Error)
	}
	chatroomDomainPtr := mapper.ChatroomStorage2Domain(*chatroom)
	return chatroomDomainPtr, nil
}
func (cp *chatroomRepo) UpdateChatroom(ctx context.Context, chatroomID chatroomDomain.ChatroomID, newRecord chatroomDomain.Chatroom) error {
	chatroom := new(types.Chatroom)

	if err := uuid.Validate(chatroomID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := newRecord.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	chatroom = mapper.ChatroomDomain2Storage(newRecord)
	chatroom.ID = chatroomID.String()

	result := cp.db.WithContext(ctx).Where("id = ?", chatroomID.String()).Updates(chatroom)
	if result.Error != nil {
		return fmt.Errorf("failed to update chatroom: %w", result.Error)
	}

	return nil
}

// message methods
func (cp *chatroomRepo) CreateMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, message chatroomDomain.Message) (chatroomDomain.MessageID, error) {
	msg := new(types.Message)
	if err := message.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("validation error: %w", err)
	}
	msg = mapper.MessageDomain2Storage(message)
	msg.ChatroomID = chatroomID.String()
	result := cp.db.WithContext(ctx).Create(msg)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create message: %w", result.Error)
	}
	id, err := uuid.Parse(msg.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create message: %w", result.Error)
	}
	return id, nil
}

func (cp *chatroomRepo) DeleteMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID) error {
	msg := new(types.Message)
	if err := uuid.Validate(messageID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	result := cp.db.WithContext(ctx).Where("id = ? AND chatroom_id = ?", messageID.String(), chatroomID.String()).Delete(msg)
	if result.Error != nil {
		return fmt.Errorf("failed to delete message: %w", result.Error)
	}
	return nil
}

func (cp *chatroomRepo) GetAllMessages(ctx context.Context, chatroomID chatroomDomain.ChatroomID) ([]chatroomDomain.Message, error) {
	var messages []types.Message
	result := cp.db.WithContext(ctx).Where("chatroom_id = ?", chatroomID.String()).Find(&messages)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get messages: %w", result.Error)
	}
	messagesDomain := fp.Map(messages, mapper.MessageStorage2Domain)
	return messagesDomain, nil
}

func (cp *chatroomRepo) GetMessageByID(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID) (*chatroomDomain.Message, error) {
	message := new(types.Message)
	if err := uuid.Validate(messageID.String()); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	result := cp.db.WithContext(ctx).Where("id = ? AND chatroom_id = ?", messageID.String(), chatroomID.String()).First(message)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get message: %w", result.Error)
	}
	messageDomainPtr := mapper.MessageStorage2Domain(*message)
	return messageDomainPtr, nil
}

func (cp *chatroomRepo) UpdateMessage(ctx context.Context, chatroomID chatroomDomain.ChatroomID, messageID chatroomDomain.MessageID, newMessage chatroomDomain.Message) error {
	oldMessage := new(types.Message)
	newMsg := new(types.Message)

	if err := uuid.Validate(messageID.String()); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := newMessage.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	result := cp.db.WithContext(ctx).Where("id = ? AND chatroom_id = ?", messageID.String(), chatroomID.String()).First(oldMessage)
	if result.Error != nil {
		return fmt.Errorf("failed to find message: %w", result.Error)
	}

	newMsg = mapper.MessageDomain2Storage(newMessage)
	newMsg.ID = messageID.String()
	newMsg.ChatroomID = chatroomID.String()

	result = cp.db.WithContext(ctx).Where("id = ? AND chatroom_id = ?", messageID.String(), chatroomID.String()).Updates(newMsg)
	if result.Error != nil {
		return fmt.Errorf("failed to update message: %w", result.Error)
	}

	return nil
}
