package types

import "time"

type Message struct {
	ID         string
	ChatroomID string
	UserID     string
	Content    string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}
