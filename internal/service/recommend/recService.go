package recommend

import (
	"context"
	"fmt"
	"strings"

	"github.com/archMqq/book-helper/internal/clients"
	"github.com/archMqq/book-helper/internal/config"
	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/models"
)

type RecService struct {
	gptClient *clients.GPTClient
}

func New(cfg *config.RecData) *RecService {
	return &RecService{
		gptClient: clients.NewGpt(cfg.GPTData),
	}
}

func (rs RecService) GetBooks(ctx context.Context, pref *models.Preferences) (string, error) {
	strPref := fmt.Sprintf("favorite genres: %s, favorite authors: %s",
		strings.Join(pref.FavoriteGenres, ", "),
		strings.Join(pref.FavoriteAuthors, ", "),
	)

	res, err := rs.gptClient.AskForNewBooks(ctx, strPref)
	if err != nil {
		return "", domain.ErrGptIsDown
	}

	return res, nil
}
