package types

type ChatRoom struct {
	Base
	Name string `gorm:"default:'main';not null;size:100"`
}
