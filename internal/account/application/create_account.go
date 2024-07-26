package application

import (
	"hc/internal/account/domain/entity"
	"hc/internal/account/domain/password"
	apiStore "hc/internal/account/domain/store"
)

type CreateAccount struct {
	Store  apiStore.Player
	Hasher password.Hasher
}

func (c *CreateAccount) Create(name, password string) (int, error) {
	username, err := entity.NewUsername(name)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := c.Hasher.Hash(password)
	if err != nil {
		return 0, err
	}

	account := entity.Account{
		Username: username,
		Password: hashedPassword,
	}

	accountId, err := c.Store.Add(account)
	if err != nil {
		return 0, err
	}

	return accountId, nil
}
