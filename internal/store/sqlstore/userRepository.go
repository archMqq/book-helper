package sqlstore

import "errors"

type UserRepository struct {
	store *Store
}

var (
	ErrUserExists = errors.New("user already exists")
)

func (ur UserRepository) Register(userID int64, userName string) error {
	query := "INSERT INTO User (TelegramID, Username) VALUES ($2, $3) ON CONFLICT DO NOTHING"

	_, err := ur.store.db.Exec(query, userID, userName)
	if err != nil {
		return err
	}

	return nil
}
