package services_test

import (
	"context"
	"testing"

	"github.com/archMqq/book-helper/internal/gpt/domain"
	"github.com/archMqq/book-helper/internal/gpt/services"
	"github.com/archMqq/book-helper/internal/gpt/services/mocks"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRecomendations(t *testing.T) {
	testCases := []struct {
		name        string
		preferences *models.Preferences
		output      string
		err         error
	}{
		{
			name: "clear",
			preferences: &models.Preferences{
				FavoriteGenres:  []string{"Фантастика"},
				FavoriteAuthors: []string{"Стругацкие"},
			},
			output: "rec",
			err:    nil,
		},
		{
			name: "gpt_error",
			preferences: &models.Preferences{
				FavoriteGenres:  []string{"Фантастика"},
				FavoriteAuthors: []string{"Стругацкие"},
			},
			output: "",
			err:    domain.ErrGptIsDown,
		},
		{
			name:        "no_preferences",
			preferences: &models.Preferences{},
			output:      "new_preferences",
			err:         nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockGPT := mocks.NewMockGPTClient(ctrl)
			mockGPT.EXPECT().
				AskForNewBooks(gomock.Any(), gomock.Any()).
				Return(tc.output, tc.err)

			recService := services.New(mockGPT)
			res, err := recService.GetBooks(context.Background(), tc.preferences)

			if tc.err != nil {
				assert.EqualError(t, tc.err, err.Error())
				assert.Empty(t, res)

			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			}
		})
	}
}
