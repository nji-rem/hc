package login

import (
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/api/session"
)

type UpdateSessionOnLoginMiddleware struct {
	SessionStore session.Store
}

func (u UpdateSessionOnLoginMiddleware) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		if err := next(sessionId, request, response); err != nil {
			return err
		}

		currentSession, err := u.SessionStore.Get(sessionId)
		if err != nil {
			return err
		}

		currentSession.Authenticated.Store(true)

		if err := u.SessionStore.Add(currentSession); err != nil {
			return err
		}

		return nil
	}
}
