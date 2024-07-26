package application

import (
	"hc/internal/account/domain/password"
	"hc/internal/account/domain/store"
)

type VerifyCredentials struct {
	Store            store.Player
	PasswordVerifier password.Verifier
}

func (v VerifyCredentials) Verify(username, password string) (bool, int, error) {
	ok, user, err := v.Store.FindByUsername(username)
	if err != nil {
		return false, 0, err
	}

	if !ok {
		return false, 0, nil
	}

	verified, err := v.PasswordVerifier.Verify([]byte(password), []byte(user.Password))
	if err != nil {
		return false, 0, err
	}

	return verified, user.ID, nil
}
