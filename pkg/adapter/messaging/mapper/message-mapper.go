package mapper

import (
	"fmt"
	wsTypes "nexu-chat/api/ws/types"
	"nexu-chat/internal/chatroom/domain"

	"github.com/google/uuid"
)

func MessageDomain2WS(msg domain.Message) wsTypes.ChatMessage {
	return wsTypes.ChatMessage{
		Type:    wsTypes.ChatMessageTypeMessage,
		Content: msg.Content,
		UserID:  msg.SenderID.String(),
		RoomID:  msg.ChatroomID.String(),
	}
}

func MessageWS2Domain(msg wsTypes.ChatMessage) domain.Message {
	userID := uuidParser(msg.UserID)
	chatroomID := uuidParser(msg.RoomID)
	return domain.Message{
		SenderID:   userID,
		ChatroomID: chatroomID,
		Content:    msg.Content,
	}
}

func uuidParser(id string) uuid.UUID {
	newID, err := uuid.Parse(id)
	if err != nil {
		fmt.Errorf("%w", err)
	}
	return newID
}
