package register

import (
	"github.com/davecgh/go-spew/spew"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/presentationlayer/parser/registration"
)

type Handler struct{}

func (h Handler) Handle(request *request.Bag, response chan<- connection.Response) error {
	registerForm, err := registration.ParseRegister(request.Body)
	if err != nil {
		return err
	}

	spew.Dump(registerForm)

	return nil
}
