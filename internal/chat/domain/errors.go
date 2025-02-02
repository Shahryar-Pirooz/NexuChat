package domain

import "errors"
var (
	ErrInvalidChatroomID = errors.New("chatroom id must not be nil")
	ErrEmptyName         = errors.New("name must not be empty")
	ErrNameTooLong       = errors.New("name must be less than 50 characters")
	ErrInvalidCreation   = errors.New("created at must not be zero time")
	ErrInvalidUpdate     = errors.New("updated at must not be zero time")
	ErrEmptyContent      = errors.New("content must not be empty")
	ErrInvalidUserID     = errors.New("user id must not be nil")
)
