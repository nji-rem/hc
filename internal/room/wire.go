package room

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	apiRoom "hc/api/room"
	"hc/internal/room/application"
	"hc/internal/room/domain"
	"hc/internal/room/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(
	ProvideCreateRoom,
	ProvideStore,
	wire.Bind(new(domain.Store), new(*store.Room)),
	wire.Bind(new(apiRoom.CreateRoom), new(*application.CreateRoom)),
)

var (
	roomStoreOnce sync.Once
	roomStore     *store.Room
)

var (
	createRoomOnce sync.Once
	createRoom     *application.CreateRoom
)

func ProvideCreateRoom(store domain.Store) *application.CreateRoom {
	createRoomOnce.Do(func() {
		createRoom = &application.CreateRoom{Store: store}
	})

	return createRoom
}

func ProvideStore(db *sqlx.DB) *store.Room {
	roomStoreOnce.Do(func() {
		roomStore = &store.Room{DB: db}
	})

	return roomStore
}
