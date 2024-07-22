package application

import (
	"errors"
	apiPassword "hc/api/account/password"
	"hc/internal/account/domain/validator"
)

func ValidatePassword(password string) (apiPassword.ValidationStatus, error) {
	err := validator.ValidatePassword(password)
	if err == nil {
		return apiPassword.Valid, nil
	}

	if errors.Is(err, validator.ErrPasswordTooShort) {
		return apiPassword.PasswordTooShort, nil
	}

	if errors.Is(err, validator.ErrPasswordTooLong) {
		return apiPassword.PasswordTooLong, nil
	}

	return apiPassword.PasswordTooLong, err
}
