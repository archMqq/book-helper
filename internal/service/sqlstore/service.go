package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/archMqq/book-helper/internal/domain"
)

type UserService struct {
	userRepository *UserRepository
}

func New(db *sql.DB) *UserService {
	return &UserService{
		userRepository: &UserRepository{
			db: db,
		},
	}
}

func (us UserService) CreateUser(userID int64, username string) error {
	err := us.userRepository.Register(userID, username)

	if strings.Contains(err.Error(), "alredy exists") {
		return domain.ErrUserExists
	}

	return nil
}
