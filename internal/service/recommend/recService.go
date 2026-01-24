package recommend

import (
	"fmt"
	"strings"

	"github.com/archMqq/book-helper/internal/clients"
	"github.com/archMqq/book-helper/internal/config"
	"github.com/archMqq/book-helper/internal/domain"
)

type RecService struct {
	gptClient *clients.GPTClient
}

func New(cfg *config.RecData) *RecService {
	return &RecService{
		gptClient: clients.NewGpt(cfg.GPTData),
	}
}

func (rs RecService) GetBooks(pref *domain.Preferences) (string, error) {
	strPref := fmt.Sprintf("favorite genres: %s, favorite authors: %s",
		strings.Join(pref.FavoriteGenres, ", "),
		strings.Join(pref.FavoriteAuthors, ", "),
	)

	res, err := rs.gptClient.AskForNewBooks(strPref)
	if err != nil {
		return "", domain.ErrGptIsDown
	}

	return res, nil
}
