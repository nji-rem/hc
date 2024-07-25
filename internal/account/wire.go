package account

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	apiAccount "hc/api/account"
	"hc/api/account/availability"
	"hc/api/account/password"
	apiStore "hc/api/account/store"
	"hc/internal/account/application"
	passwordDomain "hc/internal/account/domain/password"
	passwordInfra "hc/internal/account/infrastructure/password"
	"hc/internal/account/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(
	ProvidePlayerStore,
	wire.Bind(new(apiStore.Player), new(*store.Player)),
	ProvideCheckNameAvailabilityHandler,
	ProvidePasswordValidator,
	ProvidePasswordHasher,
	wire.Bind(new(passwordDomain.Hasher), new(*passwordInfra.HashService)),
	ProvideCreateAccountHandler,
	wire.Bind(new(apiAccount.CreateAccount), new(*application.CreateAccount)),
	ProvideVerifyCredentialsHandler,
	wire.Bind(new(apiAccount.VerifyCredentials), new(*application.VerifyCredentials)),
)

// playerStore can be a singleton; *sqlx.DB is thread-safe and is intended to be used in concurrent environments.
var playerStoreOnce sync.Once
var playerStore *store.Player

var passwordHasherOnce sync.Once
var passwordHasher *passwordInfra.HashService

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

func ProvidePasswordHasher() *passwordInfra.HashService {
	passwordHasherOnce.Do(func() {
		passwordHasher = new(passwordInfra.HashService)
	})

	return passwordHasher
}

func ProvideCreateAccountHandler(store *store.Player, hasher passwordDomain.Hasher) *application.CreateAccount {
	return &application.CreateAccount{
		Store:  store,
		Hasher: hasher,
	}
}

func ProvideVerifyCredentialsHandler(store *store.Player, hasher *passwordInfra.HashService) *application.VerifyCredentials {
	return &application.VerifyCredentials{
		Store:            store,
		PasswordVerifier: hasher,
	}
}
