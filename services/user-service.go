package services

import (
	"errors"
	"log"
	"streaming/db/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type DefaultUserService struct {
	DB *gorm.DB
}

// Constructor del servicio de usuario
func NewUserService(db *gorm.DB) *DefaultUserService {
	if db == nil {
		log.Fatal("Database connection cannot be nil")
	}
	return &DefaultUserService{DB: db}
}

// Registro de usuario con contraseña hasheada
func (s *DefaultUserService) Register(user models.User) error {
	// Verifica si ya existe un usuario con el mismo nombre de usuario o correo electrónico
	var existingUser models.User
	if err := s.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		return errors.New("username or email already exists")
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Guardar usuario en la base de datos
	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Obtener un usuario por nombre de usuario
func (s *DefaultUserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
