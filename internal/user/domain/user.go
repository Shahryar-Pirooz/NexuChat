package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrEmptyUsername = errors.New("username cannot be empty")

type UserID = uuid.UUID
type User struct {
	ID        UserID
	Username  string
	IP        string
	Connected bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FilterUser struct {
	Username  string
	IP        string
	Connected bool
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}
	return nil
}
