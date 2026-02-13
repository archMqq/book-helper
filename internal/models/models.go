package models

type User struct {
	ID   int64
	Name string
}

type Preferences struct {
	FavoriteGenres  []string
	FavoriteAuthors []string
}

type KakfaGPTAskQuery struct {
	Cmd  string
	Pref Preferences
}
