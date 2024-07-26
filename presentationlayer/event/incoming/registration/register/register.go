package register

import (
	"errors"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/event/parser/registration"
	"hc/presentationlayer/saga"
)

var ErrBodyNotFound = errors.New("body not found")

type Handler struct {
	RegistrationService saga.RegistrationService
}

func (h Handler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	registerBody, ok := request.Body.Parsed().(registration.Register)
	if !ok {
		return ErrBodyNotFound
	}

	if err := h.RegistrationService.Register(registerBody); err != nil {
		return err
	}

	return nil
}
