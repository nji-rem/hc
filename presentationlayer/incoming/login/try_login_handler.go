package login

import (
	"fmt"
	"hc/api/account"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/outgoing/login"
	"hc/presentationlayer/outgoing/message"
	"hc/presentationlayer/parser/handshake"
	"reflect"
)

type TryLoginHandler struct {
	CredentialsVerifier account.VerifyCredentials
}

func (t TryLoginHandler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	model, ok := request.Body.Parsed().(handshake.TryLogin)
	if !ok {
		return fmt.Errorf("expected type handshake.TryLogin, got %s", reflect.TypeOf(model))
	}

	validCredentials, err := t.CredentialsVerifier.Verify(model.Username, model.Password)
	if err != nil {
		return err
	}

	if validCredentials {
		response <- login.OKResponse{}
	} else {
		response <- message.ErrorResponse{Msg: "login incorrect"}
	}

	return nil
}
