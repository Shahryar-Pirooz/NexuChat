package mapper

import (
	"nexu-chat/internal/user/domain"
	"nexu-chat/pkg/adapter/storage/types"

	"github.com/google/uuid"
)

func UserDomain2Storage(src domain.User) *types.User {
	return &types.User{
		Base: types.Base{
			ID:        src.ID.String(),
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
		Username:  src.Username,
		IP:        src.IP,
		Connected: src.Connected,
	}
}

func UserStorage2Domain(src types.User) *domain.User {
	id, err := uuid.Parse(src.ID)
	if err != nil {
		return nil
	}
	return &domain.User{
		ID:        id,
		Username:  src.Username,
		IP:        src.IP,
		Connected: src.Connected,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}
