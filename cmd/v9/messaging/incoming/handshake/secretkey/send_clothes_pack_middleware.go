package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/packet"
	"hc/cmd/v9/messaging/outgoing/handshake"
	"io"
)

type SendClothesPackMiddleware struct{}

func (s SendClothesPackMiddleware) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(request *connection.Request, writer io.Writer) error {
		err := next(request, writer)
		if err != nil {
			return err
		}

		// dependency injection :)
		_, _ = handshake.NewDefaultClothesPack().WriteTo(writer)

		log.Debug().Msg("Default clothes pack sent to the client")

		return nil
	}
}
