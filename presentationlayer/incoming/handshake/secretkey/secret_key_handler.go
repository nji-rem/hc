package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
)

func HandleSecretKey(request *connection.Request, response chan<- connection.Response) error {
	log.Warn().Msg("The emulator is currently not encrypting data, secret key handler aborted")

	return nil
}
