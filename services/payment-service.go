package services

import "streaming/db/models"

type PaymentProcessor interface {
	ProcessPayment(payment models.Payment) (string, error)
}

type StripePaymentProcessor struct{}

func (s *StripePaymentProcessor) ProcessPayment(payment models.Payment) (string, error) {
	// Lógica ficticia para simular un pago con Stripe
	return "tx123456", nil
}
