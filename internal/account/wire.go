package account

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	apiStore "hc/api/account/store"
	"hc/internal/account/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(
	ProvidePlayerStore,
	wire.Bind(new(apiStore.Player), new(*store.Player)),
)

// playerStore can be a singleton; *sqlx.DB is thread-safe and is intended to be used in concurrent environments.
var playerStoreOnce sync.Once
var playerStore *store.Player

func ProvidePlayerStore(db *sqlx.DB) *store.Player {
	playerStoreOnce.Do(func() {
		playerStore = &store.Player{DB: db}
	})

	log.Info().Msg("Loaded player store")

	return playerStore
}
