package services

import (
	"errors"
	"os"
	"streaming/db/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService interface {
	Authenticate(username, password string) (string, error)
}

type DefaultAuthService struct{}

func (s *DefaultAuthService) Authenticate(username, password string) (string, error) {
	// Aquí deberías hacer la validación contra la base de datos
	user := models.User{Username: username}
	// Aquí podrías usar GORM para buscar el usuario en la base de datos
	// db.First(&user, "username = ?", username)

	// Si no existe, devolvemos un error
	if user.Username == "" {
		return "", ErrInvalidCredentials
	}

	// Compara las contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	// Generar un JWT
	token, err := generateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "myapp",
		Subject:   username,
	}

	// Aquí se genera el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
