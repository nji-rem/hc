package registration

import (
	"github.com/rs/zerolog/log"
	"hc/api/account/password"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/pkg/packet"
	"hc/presentationlayer/outgoing/registration"
)

type PasswordVerifyHandler struct {
	PasswordValidator password.ValidationFunc
}

func (p PasswordVerifyHandler) Handle(request *request.Bag, response chan<- connection.Response) error {
	reader := packet.AcquireReader(request.Body)
	defer packet.ReleaseReader(reader)

	username, err := reader.String()
	if err != nil {
		return err
	}

	password, err := reader.String()
	if err != nil {
		return err
	}

	status, err := p.PasswordValidator(password)
	if err != nil {
		return err
	}

	log.Info().Msgf("Got username %s and password %s", username, password)

	response <- registration.PasswordApprovedResponse{
		StatusCode: int(status),
	}

	return nil
}
