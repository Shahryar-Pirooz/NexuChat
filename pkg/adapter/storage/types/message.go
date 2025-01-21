package types

type Message struct {
	Base
	SenderID   string
	Sender     User `gorm:"foreignKey:SenderID"`
	ChatroomID string
	Chatroom   Chatroom `gorm:"foreignKey:ChatroomID"`
	Content    string
}
