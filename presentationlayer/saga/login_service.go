package saga

import (
	"hc/api/account"
	"hc/api/session"
	"hc/presentationlayer/sessiondata"
)

type LoginService struct {
	CredentialsVerifier account.VerifyCredentials
	SessionStore        session.Store
}

func (l LoginService) Login(sessionId, username, password string) (bool, error) {
	ok, accountId, err := l.CredentialsVerifier.Verify(username, password)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	currentSession, err := l.SessionStore.Get(sessionId)
	if err != nil {
		return false, err
	}

	currentSession.Authenticated.Store(true)
	currentSession.Set(sessiondata.AccountID, accountId)
	currentSession.Set(sessiondata.Username, username)

	if err := l.SessionStore.Add(currentSession); err != nil {
		return false, err
	}

	return true, nil
}
