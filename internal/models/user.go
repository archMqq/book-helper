package models

type User struct {
	TgID int64
	Name string
}

type Preferences struct {
	FavoriteGenres  []string
	FavoriteAuthors []string
}
