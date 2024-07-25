package handshake

import (
	"fmt"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/parser/handshake"
	"reflect"
)

type TryLoginHandler struct{}

func (t TryLoginHandler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	model, ok := request.Body.Parsed().(handshake.TryLogin)
	if !ok {
		return fmt.Errorf("expected type handshake.TryLogin, got %s", reflect.TypeOf(model))
	}

	fmt.Println(model.Username)
	fmt.Println(model.Password)

	return nil
}
