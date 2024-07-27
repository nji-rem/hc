package application

import (
	"fmt"
	"hc/api/profile"
	"hc/internal/profile/domain"
)

type UpdateProfile struct {
	Store domain.Store
}

func (u UpdateProfile) Update(accountId int, updatable profile.Updatable) error {
	foundProfile, err := u.Store.FindByAccountID(accountId)
	if err != nil {
		return fmt.Errorf("unable to find foundProfile for account id %d, err: %s", accountId, err.Error())
	}

	foundProfile.Motto = updatable.Motto

	if updatable.Figure != "" {
		foundProfile.Look = updatable.Figure
	}

	if updatable.Sex != "" {
		foundProfile.Sex = updatable.Sex
	}

	if err := u.Store.Update(foundProfile); err != nil {
		return err
	}

	return nil
}
