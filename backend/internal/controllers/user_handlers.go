package controllers

import (
    "github.com/gofiber/fiber/v2"
    "memorandum-backend/internal/usecases"  // Correct import for the usecase
    "memorandum-backend/internal/entities"
)

type UserHandler struct {
    userService *usecase.UserService  // Use the UserService from the usecase layer
}

func NewUserHandler(userService *usecase.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
    user := new(entities.User)
    
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
    }

    if user.Email == "" || user.Password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
    }

    if err := h.userService.Register(user); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    input := new(struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    })
    if err := c.BodyParser(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
    }

    token, err := h.userService.Login(input.Email, input.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    return c.JSON(fiber.Map{"token": token})
}