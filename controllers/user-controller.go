package controllers

import (
	"streaming/db/models"
	"streaming/services"

	"github.com/gofiber/fiber/v2"
)

var userService services.UserService = &services.DefaultUserService{}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user data")
	}

	if err := userService.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User registered"})
}
