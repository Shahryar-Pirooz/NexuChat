package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserID = uuid.UUID
type UserRole uint8

const (
	UserRoleUnknown UserRole = iota
	UserRoleAdmin
	UserRoleUser
)

type User struct {
	ID        UserID
	Username  string
	Password  string
	IP        string
	Role      UserRole
	Createdat time.Time
	Updatedat time.Time
}

type FilterUser struct {
	Username string
	Role     UserRole
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}

	if u.Password == "" {
		return ErrEmptyPassword
	}

	if u.IP == "" {
		return ErrEmptyIP
	}

	if u.Role == UserRoleUnknown {
		return ErrInvalidRole
	}

	if u.Createdat.IsZero() {
		return ErrInvalidCreation
	}

	if u.Updatedat.IsZero() {
		return ErrInvalidUpdate
	}

	return nil
}
