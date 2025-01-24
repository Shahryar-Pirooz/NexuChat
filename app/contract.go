package app

import (
	"context"
	"nexu-chat/config"
	chatroomPort "nexu-chat/internal/chatroom/port"
	userPort "nexu-chat/internal/user/port"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	ChatroomService(ctx context.Context) chatroomPort.Service
	NatsClient() *nats.Conn
	DB() *gorm.DB
	Config() config.Config
}
