//go:generate mockgen -source=user_service.go -destination=mocks/iser_service.go -package=mocks
package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/archMqq/book-helper/internal/bot/domain"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	Register(context.Context, int64, string) error
	GetPreferences(context.Context, int64) (*models.Preferences, error)
	SaveFavoriteAuthors(context.Context, int64, []string) error
	SaveFavoriteGenres(context.Context, int64, []string) error
}

type UserService struct {
	userRepository UserRepository
	logger         *logrus.Logger
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
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
