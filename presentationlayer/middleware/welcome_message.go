package middleware

import (
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
)

func SendWelcomeMessage(handlerFunc packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		return nil
	}
}
