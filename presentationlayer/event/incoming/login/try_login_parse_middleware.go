package login

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/event/parser/handshake"
)

func ParseTryLoginMiddleware(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		model, err := handshake.ParseTryLogin(request.Body.Raw())
		if err != nil {
			return err
		}

		request.Body.SetParsedBody(model)

		log.Debug().Msg("Parsed try login request, continue...")

		return next(sessionId, request, response)
	}
}
