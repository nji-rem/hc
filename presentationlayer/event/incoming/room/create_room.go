package room

import (
	"fmt"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/room"
	"hc/api/session"
	"hc/presentationlayer/event/viewmodel"
	"hc/presentationlayer/sessiondata"
)

type CreateRoomHandler struct {
	RoomCreator  room.CreateRoom
	SessionStore session.Store
}

func (c CreateRoomHandler) Handle(sessionId string, bag *request.Bag, response chan<- connection.Response) error {
	session, err := c.SessionStore.Get(sessionId)
	if err != nil {
		return err
	}

	accountId, ok := session.Get(sessiondata.AccountID).(int)
	if !ok {
		return fmt.Errorf("unable to retrieve account id, got %v", accountId)
	}

	roomObject, ok := bag.Body.Parsed().(viewmodel.CreateFlat)
	if !ok {
		return fmt.Errorf("unable to get viewmodel.CreateFlat, got %v instead", roomObject)
	}

	fmt.Println("room handler")

	//err := c.RoomCreator.Create(accountId, roomObject.RoomName, roomObject.RoomModel, roomObject.)

	return nil
}
