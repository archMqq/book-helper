package service

import (
	"github.com/archMqq/book-helper/internal/models"
)

type UserService interface {
	CreateUser(userID int64, username string) error
	GetPreferences(userID int64) (*models.Preferences, error)
}

type RecService interface {
	GetBooks(*models.Preferences) (string, error)
}
