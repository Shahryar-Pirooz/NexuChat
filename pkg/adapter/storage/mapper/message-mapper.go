package mapper

import (
	"nexu-chat/internal/chatroom/domain"
	"nexu-chat/pkg/adapter/storage/types"

	"github.com/google/uuid"
)

func MessageDomain2Storage(src domain.Message) *types.Message {
	return &types.Message{
		Base: types.Base{
			ID:        src.ID.String(),
			CreatedAt: src.CreateAt,
			UpdatedAt: src.UpdateAt,
		},
		SenderID:   src.SenderID.String(),
		ChatroomID: src.ChatroomID.String(),
		Content:    src.Content,
	}
}
func MessageStorage2Domain(src types.Message) *domain.Message {
	return &domain.Message{
		ID:         uuid.MustParse(src.ID),
		SenderID:   uuid.MustParse(src.SenderID),
		ChatroomID: uuid.MustParse(src.ChatroomID),
		CreateAt:   src.CreatedAt,
		UpdateAt:   src.UpdatedAt,
	}
}
