package validator

import "errors"

var (
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooLong  = errors.New("password is too long")
)

const (
	MinPasswordSize = 5
	MaxPasswordSize = 64 // TODO: Check if the UI allows passphrases
)

func ValidatePassword(password string) error {
	if len(password) < MinPasswordSize {
		return ErrPasswordTooShort
	}

	if len(password) > MaxPasswordSize {
		return ErrPasswordTooLong
	}

	return nil
}
