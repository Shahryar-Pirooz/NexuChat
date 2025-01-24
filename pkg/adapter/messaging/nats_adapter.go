package messaging

import (
	"fmt"
	chatroomDomain "nexu-chat/internal/chatroom/domain"
	"nexu-chat/internal/user/domain"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

const (
	SubjectNewMessage = "chat.message.new"
	SubjectUserJoined = "chat.user.joined"
	SubjectUserLeft   = "chat.user.left"
)

type NatsAdapter struct {
	conn *nats.Conn
}

func NewNatsAdapter(conn *nats.Conn) *NatsAdapter {
	return &NatsAdapter{conn: conn}
}

func (n *NatsAdapter) PublishMessage(msg chatroomDomain.Message) error {
	natsMsg := nats.NewMsg(SubjectNewMessage)
	natsMsg.Header.Add("user-id", msg.SenderID.String())
	natsMsg.Data = []byte(msg.Content)
	return n.conn.PublishMsg(natsMsg)
}

func (n *NatsAdapter) SubscribeMessages(handler func(chatroomDomain.Message)) error {
	_, err := n.conn.Subscribe(SubjectNewMessage, func(msg *nats.Msg) {
		userID := msg.Header.Get("user-id")
		id, err := uuid.Parse(userID)
		if err != nil {
			fmt.Println("error parsing uuid")
		}
		handler(chatroomDomain.Message{
			SenderID: domain.UserID(id),
			Content:  string(msg.Data),
		})
	})
	return err
}

func (n *NatsAdapter) PublishUserJoined(user domain.User) error {
	natsMsg := nats.NewMsg(SubjectUserJoined)
	natsMsg.Header.Add("user-id", user.ID.String())
	natsMsg.Header.Add("user-ip", user.IP)
	natsMsg.Data = []byte(user.Username)
	return n.conn.PublishMsg(natsMsg)
}

func (n *NatsAdapter) SubscribeUserJoined(handler func(domain.User)) error {
	_, err := n.conn.Subscribe(SubjectUserJoined, func(msg *nats.Msg) {
		userID := msg.Header.Get("user-id")
		id, err := uuid.Parse(userID)
		if err != nil {
			fmt.Println("error parsing uuid")
		}
		handler(domain.User{
			ID:        domain.UserID(id),
			Username:  string(msg.Data),
			IP:        msg.Header.Get("user-ip"),
			Connected: true,
		})
	})
	return err
}

func (n *NatsAdapter) PublishUserLeft(user domain.User) error {
	natsMsg := nats.NewMsg(SubjectUserLeft)
	natsMsg.Header.Add("user-id", user.ID.String())
	natsMsg.Header.Add("user-ip", user.IP)
	natsMsg.Data = []byte(user.Username)
	return n.conn.PublishMsg(natsMsg)
}

func (n *NatsAdapter) SubscribeUserLeft(handler func(domain.User)) error {
	_, err := n.conn.Subscribe(SubjectUserLeft, func(msg *nats.Msg) {
		userID := msg.Header.Get("user-id")
		id, err := uuid.Parse(userID)
		if err != nil {
			fmt.Println("error parsing uuid")
		}
		handler(domain.User{
			ID:        domain.UserID(id),
			Username:  string(msg.Data),
			IP:        msg.Header.Get("user-ip"),
			Connected: false,
		})
	})
	return err
}

func (n *NatsAdapter) Close() {
	n.conn.Close()
}
