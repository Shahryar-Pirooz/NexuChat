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
	Password  string
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

func (u *User) ComparePassword(password string) bool {
	if u.Password == password {
		return true
	}
	return false
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}
	return nil
}
