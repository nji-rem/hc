package application

import (
	"errors"
	"fmt"
	"hc/api/account/availability"
	"hc/internal/account/domain/entity"
	apiStore "hc/internal/account/domain/store"
)

type CheckNameAvailabilityHandler struct {
	Store apiStore.Player
}

func (c CheckNameAvailabilityHandler) Handle(name string) (availability.Status, error) {
	_, err := entity.NewUsername(name)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUsernameTooLong):
			return availability.UsernameTooLong, nil
		case errors.Is(err, entity.ErrUsernameTooShort):
			return availability.UsernameTooShort, nil
		case errors.Is(err, entity.ErrInvalidCharacters):
			return availability.UsernameContainsIllegalCharacters, nil
		default:
			// There's no default availability type for an actual message, so we'll just tell the client that the username
			// is currently taken. The caller can also decide to e.g. disconnect the user due to the message.
			return availability.UsernameTaken, fmt.Errorf("an unknown message occurred while checking username availability: %s", err.Error())
		}
	}

	if taken, _ := c.Store.NameTaken(entity.Username(name)); taken {
		return availability.UsernameTaken, nil
	}

	return availability.Available, nil
}
