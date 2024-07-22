package application

import (
	"errors"
	"fmt"
	"hc/api/account/availability"
	apiStore "hc/api/account/store"
	"hc/internal/account/domain/accountaggregate"
)

type CheckNameAvailabilityHandler struct {
	Store apiStore.Player
}

func (c CheckNameAvailabilityHandler) Handle(name string) (availability.Status, error) {
	_, err := accountaggregate.NewUsername(name)
	if err != nil {
		switch {
		case errors.Is(err, accountaggregate.ErrUsernameTooLong):
			return availability.UsernameTooLong, nil
		case errors.Is(err, accountaggregate.ErrUsernameTooShort):
			return availability.UsernameTooShort, nil
		case errors.Is(err, accountaggregate.ErrInvalidCharacters):
			return availability.UsernameContainsIllegalCharacters, nil
		default:
			// There's no default availability type for an actual error, so we'll just tell the client that the username
			// is currently taken. The caller can also decide to e.g. disconnect the user due to the error.
			return availability.UsernameTaken, fmt.Errorf("an unknown error occurred while checking username availability: %s", err.Error())
		}
	}

	if taken, _ := c.Store.NameTaken(accountaggregate.Username(name)); taken {
		return availability.UsernameTaken, nil
	}

	return availability.Available, nil
}
