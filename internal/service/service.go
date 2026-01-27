package service

import (
	"github.com/archMqq/book-helper/internal/models"
)

type UserService interface {
	// Возвращает ErrUserExists, если пользователь существует
	CreateUser(userID int64, username string) error

	GetPreferences(userID int64) (*models.Preferences, error)

	SaveAuthors(userID int64, authors []string) error
}

type RecService interface {
	GetBooks(*models.Preferences) (string, error)
}
