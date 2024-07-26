package middleware

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/event/parser/registration"
)

func ParseRequestMiddleware(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		parsed, err := registration.ParseRegister(request.Body.Raw())
		if err != nil {
			return err
		}

		request.Body.SetParsedBody(parsed)

		log.Debug().Msg("Request body parsed, continue registration process")

		return next(sessionId, request, response)
	}
}
