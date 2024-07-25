package application

import (
	"errors"
	apiPassword "hc/api/account/password"
	"hc/internal/account/domain/password"
)

func ValidatePassword(input string) (apiPassword.ValidationStatus, error) {
	err := password.Validate(input)
	if err == nil {
		return apiPassword.Valid, nil
	}

	if errors.Is(err, password.ErrPasswordTooShort) {
		return apiPassword.PasswordTooShort, nil
	}

	if errors.Is(err, password.ErrPasswordTooLong) {
		return apiPassword.PasswordTooLong, nil
	}

	return apiPassword.PasswordTooLong, err
}
