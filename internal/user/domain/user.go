package domain

import (
	"github.com/google/uuid"
)

type UserID = uuid.UUID
type User struct {
	ID        UserID
	Username  string
	Connected bool
}

type FilterUser struct {
	Username  string
	Connected bool
}
