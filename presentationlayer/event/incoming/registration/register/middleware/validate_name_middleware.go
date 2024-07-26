package middleware

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"hc/api/account/availability"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/event/parser/registration"
	"reflect"
)

var (
	ErrNoBodyFound          = errors.New("body is not parsed")
	ErrUsernameNotAvailable = errors.New("user tried to register a username that already exists")
)

type ValidateUsername struct {
	AvailabilityChecker availability.UsernameAvailableFunc
}

func (v ValidateUsername) Handle(next packet.HandlerFunc) packet.HandlerFunc {
	return func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
		parsedBody, ok := request.Body.Parsed().(registration.Register)
		if !ok {
			return fmt.Errorf("expected type registration.Register, got %s instead", reflect.TypeOf(parsedBody))
		}

		availabilityStatus, err := v.AvailabilityChecker(parsedBody.Username)
		if err != nil {
			return err
		}

		if availabilityStatus != availability.Available {
			return ErrUsernameNotAvailable
		}

		log.Debug().Msg("Sanity check went fine, continue registration process")

		return next(sessionId, request, response)
	}
}
