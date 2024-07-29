package middleware

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/api/session"
)

type MustBeAuthenticated struct {
	SessionStore session.Store
}

func (m MustBeAuthenticated) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		sessionBag, err := m.SessionStore.Get(sessionId)
		if err != nil {
			return err
		}

		if !sessionBag.Authenticated.Load() {
			log.Warn().Msgf("WARNING! Session %s tried to perform a call that requires authentication, but this sessionBag is unauthenticated", sessionId)

			return fmt.Errorf("sessionBag with id %s is unauthenticated", sessionId)
		}

		log.Debug().Msgf("Session ID %s passed authentication check", sessionId)

		return next(sessionId, request, response)
	}
}
