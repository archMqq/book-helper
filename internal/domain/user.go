package domain

type User struct {
	ID   int
	TgID int64
	Name string
}

type Preferences struct {
	UserID          int
	FavoriteGenres  []string
	FavoriteAuthors []string
}
