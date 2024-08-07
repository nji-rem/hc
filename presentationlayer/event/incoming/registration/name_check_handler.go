package registration

import (
	"github.com/rs/zerolog/log"
	"hc/api/account/availability"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/pkg/packet"
	"hc/presentationlayer/outgoing/registration"
)

type NameCheckHandler struct {
	availabilityChecker availability.UsernameAvailableFunc
}

func (n NameCheckHandler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	reader := packet.AcquireReader(request.Body.Raw())
	defer packet.ReleaseReader(reader)

	username, err := reader.String()
	if err != nil {
		return err
	}

	log.Info().Msgf("Received user with name %s from the client", username)

	availabilityStatus, err := n.availabilityChecker(username)
	if err != nil {
		return err
	}

	response <- registration.ApproveNameReply{
		NameCheckCode: int(availabilityStatus),
	}

	return nil
}

func NewNameCheckHandler(availabilityChecker availability.UsernameAvailableFunc) NameCheckHandler {
	return NameCheckHandler{
		availabilityChecker: availabilityChecker,
	}
}
