package password

import "errors"

var (
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooLong  = errors.New("password is too long")
)

const (
	MinPasswordSize = 6
	MaxPasswordSize = 15
)

func Validate(password string) error {
	if len(password) < MinPasswordSize {
		return ErrPasswordTooShort
	}

	if len(password) > MaxPasswordSize {
		return ErrPasswordTooLong
	}

	return nil
}
