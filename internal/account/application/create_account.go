package application

import (
	apiStore "hc/api/account/store"
	"hc/internal/account/domain/accountaggregate"
	"hc/internal/account/domain/password"
)

type CreateAccount struct {
	Store  apiStore.Player
	Hasher password.Hasher
}

func (c *CreateAccount) Create(name, password, figure, gender string) (bool, error) {
	username, err := accountaggregate.NewUsername(name)
	if err != nil {
		return false, err
	}

	hashedPassword, err := c.Hasher.Hash(password)
	if err != nil {
		return false, err
	}

	entity := accountaggregate.Entity{
		Username: username,
		Password: hashedPassword,
		Gender:   gender,
		Look:     figure,
		Motto:    "",
	}

	if err := c.Store.Add(entity); err != nil {
		return false, err
	}

	return true, nil
}
