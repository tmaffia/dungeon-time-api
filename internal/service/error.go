package service

import "errors"

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserExists        = errors.New("user already exists")
	ErrorIncorrectPassword = errors.New("incorrect password")
	ErrorInvalidUser       = errors.New("invalid user")
	ErrorInvalidRole       = errors.New("invalid role")
	ErrorInvalidTimezone   = errors.New("invalid timezone")
	ErrorInvalidPassword   = errors.New("invalid password")
	ErrorInvalidEmail      = errors.New("invalid email")
	ErrorInvalidUsername   = errors.New("invalid username")
)
