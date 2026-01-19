package sqlstore

import (
	"database/sql"

	"github.com/archMqq/book-helper/internal/store"
)

type Store struct {
	userRepository *UserRepository
	db             *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return *s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return *s.userRepository
}
