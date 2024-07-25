package application

import (
	"hc/api/account/store"
	"hc/internal/account/domain/password"
)

type VerifyCredentials struct {
	Store            store.Player
	PasswordVerifier password.Verifier
}

func (v VerifyCredentials) Verify(username, password string) (bool, error) {
	ok, user, err := v.Store.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	verified, err := v.PasswordVerifier.Verify([]byte(password), []byte(user.Password))
	if err != nil {
		return false, err
	}

	return verified, nil
}
