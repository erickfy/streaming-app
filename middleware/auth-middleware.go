package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var ErrNoAuthorizationHeader = errors.New("missing authorization header")

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": ErrNoAuthorizationHeader.Error()})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verificar el token
		claims := &jwt.RegisteredClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
		}

		// Si el token es válido, continuamos con la siguiente función
		c.Locals("user", claims.Subject)

		return c.Next()
	}
}
