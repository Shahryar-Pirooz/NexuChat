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
	return &domain.User{
		ID:        uuid.MustParse(src.ID),
		Username:  src.Username,
		IP:        src.IP,
		Connected: src.Connected,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
	}
}
