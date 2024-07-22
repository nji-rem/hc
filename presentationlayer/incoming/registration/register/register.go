package register

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/pkg/packet"
)

type Handler struct{}

func (h Handler) Handle(request *connection.Request, response chan<- connection.Response) error {
	reader := packet.AcquireReader(request.Body)
	defer packet.ReleaseReader(reader)

	username, err := reader.String()
	if err != nil {
		return err
	}

	log.Info().Msgf("Username: %s", username)

	figure, err := reader.String()
	if err != nil {
		return err
	}

	log.Info().Msgf("Figure: %s", figure)

	gender, err := reader.String()
	if err != nil {
		return err
	}

	log.Info().Msgf("Username: %s, figure: %s, gender: %s", username, figure, gender)

	return nil
}
