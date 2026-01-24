package domain

import "errors"

var (
	//UserService
	ErrUserExists = errors.New("user already exists")

	//RecService
	ErrGptIsDown = errors.New("error gpt asking")
)
