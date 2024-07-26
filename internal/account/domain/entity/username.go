package entity

import (
	"errors"
	"regexp"
)

type Username string

var (
	allowedCharactersPattern = "^[a-zA-Z0-9-@!]+$"
	allowedCharactersRegex   = regexp.MustCompile(allowedCharactersPattern)
)

var (
	ErrUsernameTooShort  = errors.New("username is too short")
	ErrUsernameTooLong   = errors.New("username is too long")
	ErrInvalidCharacters = errors.New("username contains invalid characters")
)

const (
	UsernameMinSize = 3
	UsernameMaxSize = 15
)

func NewUsername(name string) (Username, error) {
	if len(name) < UsernameMinSize {
		return "", ErrUsernameTooShort
	}

	if len(name) > UsernameMaxSize {
		return "", ErrUsernameTooLong
	}

	if !allowedCharactersRegex.MatchString(name) {
		return "", ErrInvalidCharacters
	}

	return Username(name), nil
}
