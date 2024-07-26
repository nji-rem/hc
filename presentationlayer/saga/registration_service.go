package saga

import (
	"hc/api/account"
	"hc/api/profile"
	"hc/presentationlayer/event/parser/registration"
)

type RegistrationService struct {
	// CreateAccount contains an implementation of the Identity / Authentication Service.
	CreateAccount account.CreateAccount

	// CreateProfile contains an implementation of the Profile Service.
	CreateProfile profile.CreateProfile
}

func (r RegistrationService) Register(register registration.Register) error {
	accountId, err := r.CreateAccount.Create(register.Username, register.Password)
	if err != nil {
		return err
	}

	if err := r.CreateProfile.Create(accountId, "", register.Figure, register.Sex); err != nil {
		return err
	}

	return nil
}
