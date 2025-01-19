package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrEmptyUsername = errors.New("username cannot be empty")

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

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}
	return nil
}
