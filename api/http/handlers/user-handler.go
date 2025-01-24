package handlers

import (
	"nexu-chat/internal/user/domain"
	"nexu-chat/internal/user/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.Service
}

func NewUserHandler(userService port.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register handles user registration
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		IP       string `json:"ip"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
		IP:       c.IP(),
	}

	createdUser, err := h.userService.CreateUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": createdUser,
	})
}

// Login handles user login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.AuthenticateUser(c.Context(), request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
	return c.JSON(fiber.Map{
		"user": user,
	})
}
