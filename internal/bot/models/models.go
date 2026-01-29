package models

type User struct {
	ID   int64
	Name string
}

type Preferences struct {
	FavoriteGenres  []string // JSONB
	FavoriteAuthors []string // JSONB
}
