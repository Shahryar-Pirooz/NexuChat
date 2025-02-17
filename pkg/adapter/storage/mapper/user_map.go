package mapper

import (
	userDomain "nexu-chat/internal/user/domain"
	"nexu-chat/pkg/adapter/storage/types"

	"github.com/google/uuid"
)

func UserDomain2Storage(src userDomain.User) types.User {
	return types.User{
		ID:        src.ID.String(),
		Username:  src.Username,
		Password:  src.Password,
		IP:        src.IP,
		Role:      uint8(src.Role),
		Createdat: src.Createdat,
		Updatedat: src.Updatedat,
	}
}
func UserStorage2SDomain(src types.User) userDomain.User {
	id, _ := uuid.Parse(src.ID)
	return userDomain.User{
		ID:        id,
		Username:  src.Username,
		Password:  src.Password,
		IP:        src.IP,
		Role:      userDomain.UserRole(src.Role),
		Createdat: src.Createdat,
		Updatedat: src.Updatedat,
	}
}
