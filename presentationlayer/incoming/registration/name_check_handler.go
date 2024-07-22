package registration

import (
	"github.com/davecgh/go-spew/spew"
	"hc/api/account/availability"
	"hc/api/connection"
	"io"
)

type NameCheckHandler struct {
	availabilityChecker availability.UsernameAvailableFunc
}

func (n NameCheckHandler) Handle(request *connection.Request, writer io.Writer) error {
	spew.Dump(request)

	return nil
}

func NewNameCheckHandler(availabilityChecker availability.UsernameAvailableFunc) NameCheckHandler {
	return NameCheckHandler{availabilityChecker: availabilityChecker}
}
