package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/archMqq/book-helper/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepository: repository.NewUser(db),
	}
}

func (us UserService) CreateUser(userID int64, username string) error {
	err := us.userRepository.Register(userID, username)

	if strings.Contains(err.Error(), "alredy exists") {
		return domain.ErrUserExists
	}

	return nil
}

func (us UserService) GetPreferences(userID int64) (*models.Preferences, error) {
	res, err := us.userRepository.GetPreferences(userID)
	if err != nil {
		return nil, domain.ErrDatabaseRequest
	}

	return res, nil
}
