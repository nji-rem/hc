package secretkey

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/outgoing/handshake"
)

type SendClothesPackMiddleware struct{}

func (s SendClothesPackMiddleware) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		err := next(sessionId, request, response)
		if err != nil {
			return err
		}

		// dependency injection :)
		response <- handshake.NewDefaultClothesPack()

		log.Debug().Msg("Default clothes pack sent to the client")

		return nil
	}
}
