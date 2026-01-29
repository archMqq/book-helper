package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/archMqq/book-helper/internal/bot/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) Register(ctx context.Context, userID int64, username string) error {
	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = ur.registerUser(ctx, userID, username)
	if err != nil {
		return err
	}

	err = ur.registerUserPrefernces(ctx, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ur UserRepository) registerUser(ctx context.Context, userID int64, username string) error {
	query := "INSERT INTO users (id, username, creation_time) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"

	res, err := ur.db.ExecContext(ctx, query, userID, username, time.Now())
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("user already exists %w", err)
	}

	return nil
}

func (ur UserRepository) registerUserPrefernces(ctx context.Context, userID int64) error {
	query := "INSERT INTO preferences (user_id) VALUES ($1)"

	_, err := ur.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) GetPreferences(ctx context.Context, userID int64) (*models.Preferences, error) {
	query := "SELECT favorite_genres, favorite_authors FROM preferences WHERE user_id = $1"

	rows, err := ur.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pref := models.Preferences{}

	for rows.Next() {
		var genres, authors string
		rows.Scan(&genres, &authors)

		err = json.Unmarshal([]byte(genres), &pref.FavoriteGenres)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(authors), &pref.FavoriteAuthors)
		if err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pref, nil
}

func (ur UserRepository) SaveFavoriteAuthors(ctx context.Context, userID int64, authors []string) error {
	authorsJSON, err := json.Marshal(authors)
	if err != nil {
		return fmt.Errorf("error slice marshalling: %w", err)
	}

	query := "UPDATE preferences SET favorite_authors = $1 WHERE user_id = $2"

	_, err = ur.db.ExecContext(ctx, query, authorsJSON, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) SaveFavoriteGenres(ctx context.Context, userID int64, genres []string) error {
	genresJSON, err := json.Marshal(genres)
	if err != nil {
		return fmt.Errorf("error slice marshalling: %w", err)
	}

	query := "UPDATE preferences SET favorite_genres = $1 WHERE user_id = $2"

	_, err = ur.db.ExecContext(ctx, query, genresJSON, userID)
	if err != nil {
		return err
	}

	return nil
}
