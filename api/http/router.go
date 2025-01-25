package http

import (
	httpHandler "nexu-chat/api/http/handler"
	wsHandler "nexu-chat/api/ws/handler"
	chatroomPort "nexu-chat/internal/chatroom/port"
	userPort "nexu-chat/internal/user/port"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewRouter(
	wsHandler *wsHandler.Handler,
	userService userPort.Service,
	chatroomService chatroomPort.Service,
) *fiber.App {
	app := fiber.New()

	// Create handlers
	userHandler := httpHandler.NewUserHandler(userService)
	// chatroomHandler := handler.NewChatroomHandler(chatroomService)

	// API routes
	api := app.Group("/api")

	// User routes
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Chatroom routes
	// api.Post("/chatrooms", chatroomHandler.Create)
	// api.Get("/chatrooms", chatroomHandler.List)
	// api.Get("/chatrooms/:id", chatroomHandler.Get)

	// WebSocket route
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(wsHandler.HandleWebSocket))

	return app
}
