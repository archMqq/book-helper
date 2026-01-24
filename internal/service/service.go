package service

import "github.com/archMqq/book-helper/internal/domain"

type UserService interface {
	CreateUser(userID int64, username string) error
}

type RecService interface {
	GetBooks(*domain.Preferences) (string, error)
}
