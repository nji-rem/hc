package room

import (
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	roomParser "hc/presentationlayer/event/parser/room"
)

func ParseRoomObjectMiddleware(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		r, err := roomParser.ParseCreateFlat(request.Body.Raw())
		if err != nil {
			return err
		}

		request.Body.SetParsedBody(r)

		return next(sessionId, request, response)
	}
}
