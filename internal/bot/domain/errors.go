package domain

import "errors"

var (
	//UserService
	ErrUserExists      = errors.New("user already exists")
	ErrDatabaseRequest = errors.New("database is down")
)
