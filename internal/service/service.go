package service

import (
	"context"

	"github.com/archMqq/book-helper/internal/models"
)

type UserService interface {
	// Возвращает ErrUserExists, если пользователь существует
	CreateUser(context.Context, int64, string) error

	GetPreferences(context.Context, int64) (*models.Preferences, error)

	SaveAuthors(context.Context, int64, []string) error
	SaveGenres(context.Context, int64, []string) error
}

type RecService interface {
	GetBooks(context.Context, *models.Preferences) (string, error)
}
