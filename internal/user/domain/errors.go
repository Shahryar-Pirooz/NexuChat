package domain

import "errors"

var (
	ErrEmptyUsername   = errors.New("username cannot be empty")
	ErrEmptyPassword   = errors.New("password cannot be empty")
	ErrEmptyIP         = errors.New("IP address cannot be empty")
	ErrInvalidRole     = errors.New("invalid role provided")
	ErrInvalidCreation = errors.New("error occurred during creation")
	ErrInvalidUpdate   = errors.New("error occurred during update")
)
