//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package recommend

import (
	"context"
	"fmt"
	"strings"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/models"
)

type GPTClient interface {
	AskForNewBooks(context.Context, string) (string, error)
}

type RecService struct {
	gptClient GPTClient
}

func New(gptClient GPTClient) *RecService {
	return &RecService{
		gptClient: gptClient,
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
