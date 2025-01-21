package types

type Chatroom struct {
	Base
	Name string `gorm:"default:'main';not null;size:100"`
}
