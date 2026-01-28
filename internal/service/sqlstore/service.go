package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
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

func (us UserService) CreateUser(ctx context.Context, userID int64, username string) error {
	err := us.userRepository.Register(ctx, userID, username)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return domain.ErrUserExists
		}

		return err
	}

	return nil
}

func (us UserService) GetPreferences(ctx context.Context, userID int64) (*models.Preferences, error) {
	res, err := us.userRepository.GetPreferences(ctx, userID)
	if err != nil {
		fmt.Println(err)
		return nil, domain.ErrDatabaseRequest
	}

	return res, nil
}

func (us UserService) SaveAuthors(ctx context.Context, userID int64, authors []string) error {
	err := us.userRepository.SaveFavoriteAuthors(ctx, userID, authors)
	if err != nil {
		if strings.Contains(err.Error(), "marshalling") {
			return err
		}
		return fmt.Errorf("%w: %w", domain.ErrDatabaseRequest, err)
	}

	return nil
}

func (us UserService) SaveGenres(ctx context.Context, userID int64, genres []string) error {
	err := us.userRepository.SaveFavoriteGenres(ctx, userID, genres)
	if err != nil {
		if strings.Contains(err.Error(), "marshalling") {
			return err
		}
		return fmt.Errorf("%w: %w", domain.ErrDatabaseRequest, err)
	}

	return nil
}
