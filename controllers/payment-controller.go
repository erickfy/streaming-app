package controllers

import (
	"streaming/db/models"
	"streaming/services"

	"github.com/gofiber/fiber/v2"
)

var paymentService services.PaymentProcessor = &services.StripePaymentProcessor{}

func ProcessPayment(c *fiber.Ctx) error {
	var payment models.Payment
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid payment data")
	}

	txID, err := paymentService.ProcessPayment(payment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Payment processing failed")
	}

	return c.JSON(fiber.Map{"transactionID": txID, "status": "success"})
}
