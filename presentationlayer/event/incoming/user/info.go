package user

import (
	"errors"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/profile"
	"hc/api/session"
	"hc/presentationlayer/outgoing/user"
	"hc/presentationlayer/sessiondata"
)

var ErrAccountIdNotFound = errors.New("account id not found in session")

type InfoHandler struct {
	SessionStore  session.Store
	InfoRetriever profile.InfoRetriever
}

func (i InfoHandler) Handle(sessionId string, bag *request.Bag, response chan<- connection.Response) error {
	currentSession, err := i.SessionStore.Get(sessionId)
	if err != nil {
		return err
	}

	accountId, ok := currentSession.Get(sessiondata.AccountID).(int)
	if !ok {
		return ErrAccountIdNotFound
	}

	userInfo, err := i.InfoRetriever.Retrieve(accountId)
	if err != nil {
		return err
	}

	response <- user.InfoResponse{
		Name:   currentSession.Get(sessiondata.Username).(string),
		Figure: userInfo.Figure,
		Sex:    userInfo.Sex,
		Motto:  userInfo.Motto,
	}

	return nil
}
