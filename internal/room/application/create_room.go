package application

import (
	"github.com/rs/zerolog/log"
	"hc/api/room"
	"hc/internal/room/domain"
)

type CreateRoom struct {
	Store domain.Store
}

func (c CreateRoom) Create(accountId int, name, model, accessType string, roomOwnerVisible bool) error {
	roomName, err := domain.NewRoomName(name)
	if err != nil {
		return room.ErrInvalidRoomName
	}

	roomModel, err := domain.NewRoomModel(model)
	if err != nil {
		return room.ErrInvalidRoomModel
	}

	r := domain.Room{
		AccountID:        accountId,
		Name:             roomName,
		Model:            roomModel,
		Description:      "",
		RoomAccessType:   domain.NewRoomAccessType(accessType),
		RoomOwnerVisible: roomOwnerVisible,
	}

	if err := c.Store.Add(r); err != nil {
		return err
	}

	log.Debug().Msgf("Room '%s' successfully created", roomName)

	return nil
}
