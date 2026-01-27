package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/archMqq/book-helper/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepository *repository.UserRepository
	logger         *logrus.Logger
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
	} else if err != nil {
		return err
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

func (us UserService) SaveAuthors(userID int64, authors []string) error {
	err := us.userRepository.SaveFavoriteAuthors(userID, authors)
	if err != nil {
		return domain.ErrDatabaseRequest
	}

	return nil
}
