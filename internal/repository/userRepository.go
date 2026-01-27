package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/archMqq/book-helper/internal/models"
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
	err := ur.registerUser(userID, username)
	if err != nil {
		return err
	}

	err = ur.registerUserPrefernces(userID)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) registerUser(userID int64, username string) error {
	query := "INSERT INTO users (id, username, creation_time) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"

	res, err := ur.db.Exec(query, userID, username, time.Now)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("user already exists %w", err)
	}

	return nil
}

func (ur UserRepository) registerUserPrefernces(userID int64) error {
	query := "INSERT INTO preferences (user_id) VALUES ($1)"

	_, err := ur.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) GetPreferences(userID int64) (*models.Preferences, error) {
	query := "GET favorite_grenres, favorite_authors FROM preferences WHERE user_id = $1"

	rows, err := ur.db.Query(query, userID)
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

	return &pref, nil
}

func (ur UserRepository) SaveFavoriteAuthors(userID int64, authors []string) error {
	authorsJSON, err := json.Marshal(authors)
	if err != nil {
		return fmt.Errorf("error slice marshalling: %s", err)
	}

	query := "INSERT INTO preferences (favorite_authors) VALUES ($1) WHERE user_id = $2"

	_, err = ur.db.Exec(query, authorsJSON, userID)
	if err != nil {
		return fmt.Errorf("error user's preferences insert")
	}

	return nil
}
