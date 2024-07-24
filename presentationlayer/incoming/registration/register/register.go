package register

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
)

type Handler struct{}

func (h Handler) Handle(sessionId string, request *request.Bag, response chan<- connection.Response) error {
	spew.Dump(request.Body.Body())

	log.Info().Msg("yay registration here")

	return nil
}
