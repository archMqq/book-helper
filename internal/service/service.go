package service

type UserService interface {
	CreateUser(userID int64, username string) error
}
