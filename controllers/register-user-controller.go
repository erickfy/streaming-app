package controllers

import (
	"streaming/db/models"
	"streaming/services"

	"github.com/gofiber/fiber/v2"
)

type RegisterController struct {
	UserService services.UserService
}

// Constructor para RegisterController
func NewRegisterController(userService services.UserService) *RegisterController {
	return &RegisterController{UserService: userService}
}

func (rc *RegisterController) RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user data")
	}

	if err := rc.UserService.Register(user); err != nil {
		if err.Error() == "username or email already exists" {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User registered"})
}
