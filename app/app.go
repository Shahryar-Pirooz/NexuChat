package app

import (
	"context"
	"fmt"
	"nexu-chat/config"
	"nexu-chat/internal/chatroom"
	chatroomPort "nexu-chat/internal/chatroom/port"
	"nexu-chat/internal/user"
	userPort "nexu-chat/internal/user/port"
	"nexu-chat/pkg/adapter/storage"
	appcontext "nexu-chat/pkg/app-context"
	"nexu-chat/pkg/postgres"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type app struct {
	db              *gorm.DB
	config          config.Config
	natsClient      *nats.Conn
	userService     userPort.Service
	chatroomService chatroomPort.Service
}

// Database Part
func (a *app) setDB() error {
	db, err := postgres.NewDB(&postgres.DBOptions{
		Host:     a.config.Database.Host,
		User:     a.config.Database.User,
		Password: a.config.Database.Password,
		DBName:   a.config.Database.DBName,
		SSLMode:  a.config.Database.SSLMode,
	})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func (a *app) DB() *gorm.DB {
	return a.db
}

// Nats Part
func (a *app) setNats() error {
	URL := fmt.Sprintf("%s:%s", a.config.Nats.Host, a.config.Nats.Port)
	n, err := nats.Connect(URL)
	if err != nil {
		return err
	}
	a.natsClient = n
	return nil
}

func (a *app) NatsClient() *nats.Conn {
	return a.natsClient
}

func (a *app) Cleanup() error {
	if a.natsClient != nil {
		a.natsClient.Close()
	}
	return nil
}

// Services Part
func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appcontext.GetDatabase(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}
	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}

func (a *app) ChatroomService(ctx context.Context) chatroomPort.Service {
	db := appcontext.GetDatabase(ctx)
	if db == nil {
		if a.chatroomService == nil {
			a.chatroomService = a.chatroomServiceWithDB(a.db)
		}
		return a.chatroomService
	}
	return a.chatroomServiceWithDB(db)
}

func (a *app) chatroomServiceWithDB(db *gorm.DB) chatroomPort.Service {
	return chatroom.NewService(storage.NewChatroomRepo(db))
}

// Config Part
func (a *app) Config() config.Config {
	return a.config
}

// Application Part
func NewApp(cnfg config.Config) (App, error) {
	a := &app{
		config: cnfg,
	}
	if err := a.setDB(); err != nil {
		return nil, err
	}
	if err := a.setNats(); err != nil {
		return nil, err
	}
	return a, nil
}
