package mapper

import (
	chatDomain "nexu-chat/internal/chat/domain"
	"nexu-chat/pkg/adapter/storage/types"

	"github.com/google/uuid"
)

func ChatroomDomain2Storage(src chatDomain.Chatroom) types.Chatroom {
	return types.Chatroom{
		ID:        src.ID.String(),
		Name:      src.Name,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

func ChatroomStorage2Domain(src types.Chatroom) chatDomain.Chatroom {
	id, _ := uuid.Parse(src.ID)
	return chatDomain.Chatroom{
		ID:        id,
		Name:      src.Name,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}

func MessageDomain2Storage(src chatDomain.Message) types.Message {
	return types.Message{
		ID:         src.ID.String(),
		ChatroomID: src.ChatroomID.String(),
		UserID:     src.UserID.String(),
		Content:    src.Content,
		CreatedAt:  src.CreatedAt,
		UpdateAt:   src.UpdateAt,
	}
}

func MessageStorage2Domain(src types.Message) chatDomain.Message {
	id, _ := uuid.Parse(src.ID)
	chatroomID, _ := uuid.Parse(src.ChatroomID)
	userID, _ := uuid.Parse(src.UserID)
	return chatDomain.Message{
		ID:         id,
		ChatroomID: chatroomID,
		UserID:     userID,
		Content:    src.Content,
		CreatedAt:  src.CreatedAt,
		UpdateAt:   src.UpdateAt,
	}
}
