package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"io"
)

func HandleSecretKey(request *connection.Request, writer io.Writer) error {
	log.Warn().Msg("The emulator is currently not encrypting data, secret key handler aborted")

	return nil
}
