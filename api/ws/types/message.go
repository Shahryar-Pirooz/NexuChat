package types

type ChatMessageType uint8

const (
	ChatMessageTypeUnkonwn ChatMessageType = iota
	ChatMessageTypeMessage
	ChatMessageTypeJoin
	ChatMessageTypeLeave
)

type ChatMessage struct {
	Type    ChatMessageType `json:"type"`
	Content string          `json:"content"`
	UserID  string          `json:"user_id"`
	RoomID  string          `json:"room_id"`
}
