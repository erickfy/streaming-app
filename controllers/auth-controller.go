package controllers

import (
	"streaming/db/models"
	"streaming/services"

	"github.com/gofiber/fiber/v2"
)

var authService services.AuthService = &services.DefaultAuthService{}

func Login(c *fiber.Ctx) error {
	var creds models.User
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	token, err := authService.Authenticate(creds.Username, creds.Password)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}
