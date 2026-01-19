package store

type UserRepository interface {
	Register(userID int64, userName string) error
}
