package mapper

import (
	"nexu-chat/internal/chatroom/domain"
	"nexu-chat/pkg/adapter/storage/types"

	"github.com/google/uuid"
)

func ChatroomDomain2Storage(src domain.Chatroom) *types.Chatroom {
	return &types.Chatroom{
		Base: types.Base{
			ID:        src.ID.String(),
			CreatedAt: src.CreateAt,
			UpdatedAt: src.UpdateAt,
		},
		Name: src.Name,
	}
}
func ChatroomStorage2Domain(src types.Chatroom) *domain.Chatroom {
	return &domain.Chatroom{
		ID:       uuid.MustParse(src.ID),
		Name:     src.Name,
		CreateAt: src.CreatedAt,
		UpdateAt: src.UpdatedAt,
	}
}
