package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/packet"
	"hc/presentationlayer/outgoing/handshake"
)

// SessionDataMiddleware is responsible for sending session data to the client after the secret key handler is finished
// with its work.
type SessionDataMiddleware struct{}

func (s SessionDataMiddleware) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(request *connection.Request, response chan<- connection.Response) error {
		err := next(request, response)
		if err != nil {
			return err
		}

		// dependency injection :)
		response <- handshake.NewSessionParametersComposer()

		log.Debug().Msg("Session data sent to the client")

		return nil
	}
}
