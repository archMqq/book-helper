package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

var (
	ErrUserExists = errors.New("user already exists")
)

func (ur UserRepository) Register(userID int64, username string) error {
	query := "INSERT INTO User (TelegramID, Username) VALUES ($2, $3) ON CONFLICT DO NOTHING"

	res, err := ur.db.Exec(query, userID, username)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("user already exists %w", err)
	}

	return nil
}
