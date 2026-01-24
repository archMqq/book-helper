package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) Register(userID int64, username string) error {
	query := "INSERT INTO User (TelegramID, Username, CreateTime) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"

	res, err := ur.db.Exec(query, userID, username, time.Now)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("user already exists %w", err)
	}

	return nil
}
