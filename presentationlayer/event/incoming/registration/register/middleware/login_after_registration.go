package middleware

import (
	"fmt"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/event/parser/registration"
	"hc/presentationlayer/outgoing/login"
	"hc/presentationlayer/saga"
	"reflect"
)

type LoginAfterRegistration struct {
	LoginService saga.LoginService
}

func (l LoginAfterRegistration) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		if err := next(sessionId, request, response); err != nil {
			return err
		}

		registerBody, ok := request.Body.Parsed().(registration.Register)
		if !ok {
			return fmt.Errorf("unable to get parsed registration.Register object, got %s", reflect.TypeOf(request.Body.Parsed()))
		}

		if _, err := l.LoginService.Login(sessionId, registerBody.Username, registerBody.Password); err != nil {
			return err
		}

		response <- login.OKResponse{}

		return nil
	}
}
