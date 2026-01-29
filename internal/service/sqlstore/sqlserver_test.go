package sqlstore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/service/sqlstore"
	"github.com/archMqq/book-helper/internal/service/sqlstore/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	testCases := []struct {
		name     string
		id       int64
		username string
		returnal error
	}{
		{
			name:     "no_error",
			id:       int64(123),
			username: "Ivan",
			returnal: nil,
		},
		{
			name:     "not_new_user",
			id:       int64(123),
			username: "Ivan",
			returnal: domain.ErrUserExists,
		},
		{
			name:     "db_err",
			id:       int64(123),
			username: "Ivan",
			returnal: errors.New("unknown"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				Register(gomock.Any(), tc.id, tc.username).
				Return(tc.returnal)

			srv := sqlstore.NewUserService(mockRepo)
			err := srv.CreateUser(context.Background(), tc.id, tc.username)

			if tc.returnal == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.returnal, err)
			}
		})
	}
}
