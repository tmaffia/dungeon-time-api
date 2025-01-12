package service

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserExists        = errors.New("user already exists")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrInvalidUser       = errors.New("invalid user")
	ErrInvalidRole       = errors.New("invalid role")
	ErrInvalidTimezone   = errors.New("invalid timezone")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidUsername   = errors.New("invalid username")
)
