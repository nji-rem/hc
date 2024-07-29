package application

import (
	"hc/api/room"
	"hc/internal/room/domain"
)

type CreateRoom struct {
	Store domain.Store
}

func (c CreateRoom) Create(accountId int, name, model, description, accessType string, roomOwnerVisible bool) error {
	roomName, err := domain.NewRoomName(name)
	if err != nil {
		return room.ErrInvalidRoomName
	}

	roomModel, err := domain.NewRoomModel(model)
	if err != nil {
		return room.ErrInvalidRoomModel
	}

	r := domain.Room{
		Name:             roomName,
		Model:            roomModel,
		Description:      description,
		RoomAccessType:   domain.NewRoomAccessType(accessType),
		RoomOwnerVisible: roomOwnerVisible,
	}

	if err := c.Store.Add(r); err != nil {
		return err
	}

	return nil
}
