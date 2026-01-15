package repository

import "github.com/archMqq/book-helper/internal/models"

type UserRepository struct {
}

func (ur UserRepository) Register(userID int64, userName string) *models.User {
	return nil
}
