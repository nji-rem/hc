package application

import (
	"hc/internal/profile/domain"
)

type CreateProfile struct {
	Store domain.Store
}

func (c CreateProfile) Create(accountID int, motto, figure, sex string) error {
	accountEntity := domain.Profile{
		AccountID: accountID,
		Look:      figure,
		Motto:     motto,
		Sex:       sex,
	}

	if err := c.Store.Add(accountEntity); err != nil {
		return err
	}

	return nil
}
