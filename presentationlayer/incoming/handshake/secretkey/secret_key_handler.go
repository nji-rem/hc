package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
)

func HandleSecretKey(request *request.Bag, response chan<- connection.Response) error {
	log.Warn().Msg("The emulator is currently not encrypting data, secret key handler aborted")

	return nil
}
