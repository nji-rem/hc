package room

import (
	"github.com/davecgh/go-spew/spew"
	"hc/api/connection"
	"hc/api/connection/request"
)

type CreateRoomHandler struct {
}

func (c CreateRoomHandler) Handle(sessionId string, bag *request.Bag, response chan<- connection.Response) error {
	spew.Dump(bag.Body.Parsed())
	return nil
}
