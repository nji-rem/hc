package register

import (
	"errors"
	"hc/api/account"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/parser/registration"
)

var ErrBodyNotFound = errors.New("body not found")

type Handler struct {
	AccountCreator account.CreateAccount
}

func (h Handler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	body, ok := request.Body.Body()
	if !ok {
		return ErrBodyNotFound
	}

	registerBody, ok := body.(registration.Register)
	if !ok {
		return ErrBodyNotFound
	}

	_, err := h.AccountCreator.Create(registerBody.Username, registerBody.Password, registerBody.Figure, registerBody.Sex)
	if err != nil {
		return err
	}

	return nil
}
