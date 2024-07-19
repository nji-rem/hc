package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/packet"
	"hc/messaging/outgoing/handshake"
	"io"
)

// SessionDataMiddleware is responsible for sending session data to the client after the secret key handler is finished
// with its work.
type SessionDataMiddleware struct{}

func (s SessionDataMiddleware) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(request *connection.Request, writer io.Writer) error {
		err := next(request, writer)
		if err != nil {
			return err
		}

		// dependency injection :)
		_, _ = handshake.NewSessionParametersComposer().WriteTo(writer)

		log.Debug().Msg("Session data sent to the client")

		return nil
	}
}
