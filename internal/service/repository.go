package service

type UserRepository interface {
	Register(userID int64, userName string) error
}
