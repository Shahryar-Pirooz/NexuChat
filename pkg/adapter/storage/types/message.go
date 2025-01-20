package types

type Message struct {
	Base
	SenderID   string
	Sender     User `gorm:"foreignKey:SenderID"`
	ChatroomID string
	ChatRoom   ChatRoom `gorm:"foreignKey:ChatroomID"`
	Content    string
}
