package login

import (
	"fmt"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/event/parser/handshake"
	"hc/presentationlayer/outgoing/login"
	"hc/presentationlayer/outgoing/message"
	"hc/presentationlayer/saga"
	"reflect"
)

type TryLoginHandler struct {
	LoginService saga.LoginService
}

func (t TryLoginHandler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	model, ok := request.Body.Parsed().(handshake.TryLogin)
	if !ok {
		return fmt.Errorf("expected type handshake.TryLogin, got %s", reflect.TypeOf(model))
	}

	loginSuccess, err := t.LoginService.Login(sessionId, model.Username, model.Password)
	if err != nil {
		return err
	}

	if !loginSuccess {
		response <- message.ErrorResponse{Msg: "login incorrect"}
		return nil
	}

	response <- login.OKResponse{}

	return nil
}
