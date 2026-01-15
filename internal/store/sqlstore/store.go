package sqlstore

import "github.com/archMqq/book-helper/internal/repository"

type Store struct {
	userRepository *repository.UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() *repository.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &repository.UserRepository{}
	return s.userRepository
}
