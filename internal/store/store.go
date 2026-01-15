package store

import "github.com/archMqq/book-helper/internal/repository"

type Store interface {
	User() *repository.UserRepository
}
