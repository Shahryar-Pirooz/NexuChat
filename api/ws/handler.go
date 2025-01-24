package ws

import (
	"nexu-chat/api/ws/types"
	"nexu-chat/pkg/adapter/messaging"
	"nexu-chat/pkg/adapter/messaging/mapper"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Handler struct {
	nats    *messaging.NatsAdapter
	clients sync.Map
}

func NewHandler(nats *messaging.NatsAdapter) *Handler {
	return &Handler{
		nats: nats,
	}
}

func (h *Handler) HandleWebSocket(c *websocket.Conn) {
	h.clients.Store(c, true)
	defer func() {
		h.clients.Delete(c)
		c.Close()
	}()

	for {
		var msg types.ChatMessage
		err := c.ReadJSON(&msg)
		if err != nil {
			break
		}
		newMsg := mapper.MessageWS2Domain(msg)
		h.nats.PublishMessage(newMsg)
	}
}

func (h *Handler) broadcastMessage(message types.ChatMessage) {
	h.clients.Range(func(key, value interface{}) bool {
		conn := key.(*websocket.Conn)
		conn.WriteJSON(message)
		return true
	})
}
