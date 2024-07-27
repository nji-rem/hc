package user

import (
	"fmt"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/profile"
	"hc/api/session"
	"hc/presentationlayer/event/parser/registration"
	"hc/presentationlayer/outgoing/user"
	"hc/presentationlayer/sessiondata"
	"reflect"
)

type Update struct {
	ProfileUpdater profile.Updater
	SessionStore   session.Store
}

func (u Update) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	model, ok := request.Body.Parsed().(registration.Register)
	if !ok {
		return fmt.Errorf("expected registration.Register model, got %s", reflect.TypeOf(model))
	}

	session, err := u.SessionStore.Get(sessionId)
	if err != nil {
		return err
	}

	accountId, ok := session.Get(sessiondata.AccountID).(int)
	if !ok {
		return fmt.Errorf("expected account id to be an integer got %s", reflect.TypeOf(accountId))
	}

	if err := u.ProfileUpdater.Update(accountId, profile.Updatable{
		Motto:  model.CustomData,
		Figure: model.Figure,
		Sex:    model.Sex,
	}); err != nil {
		return err
	}

	response <- user.InfoResponse{
		Name:   session.Get(sessiondata.Username).(string),
		Figure: model.Figure,
		Sex:    model.Sex,
		Motto:  model.CustomData,
	}

	return nil
}
