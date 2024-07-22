package account

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"hc/api/account/availability"
	"hc/api/account/password"
	apiStore "hc/api/account/store"
	"hc/internal/account/application"
	"hc/internal/account/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(
	ProvidePlayerStore,
	wire.Bind(new(apiStore.Player), new(*store.Player)),
	ProvideCheckNameAvailabilityHandler,
	ProvidePasswordValidator,
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

func ProvideCheckNameAvailabilityHandler(store apiStore.Player) availability.UsernameAvailableFunc {
	handler := application.CheckNameAvailabilityHandler{Store: store}

	return handler.Handle
}

func ProvidePasswordValidator() password.ValidationFunc {
	return application.ValidatePassword
}
