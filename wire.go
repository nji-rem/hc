//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	apiAccount "hc/api/account"
	"hc/api/account/availability"
	"hc/api/account/password"
	apiConfig "hc/api/config"
	apiPacket "hc/api/packet"
	apiProfile "hc/api/profile"
	"hc/internal/account"
	"hc/internal/connection"
	"hc/internal/packet"
	"hc/internal/profile"
	"hc/internal/profile/application"
	"hc/internal/session"
	"hc/pkg/config"
	"hc/pkg/database"
	"hc/presentationlayer/event/incoming/login"
	registration2 "hc/presentationlayer/event/incoming/registration"
	"hc/presentationlayer/event/incoming/registration/register"
	"hc/presentationlayer/event/incoming/registration/register/middleware"
	"hc/presentationlayer/event/incoming/user"
	"hc/presentationlayer/saga"
	"strconv"
	"sync"
)

var ProfileSet = wire.NewSet(
	profile.Set,
	wire.Bind(new(apiProfile.CreateProfile), new(*application.CreateProfile)),
	wire.Bind(new(apiProfile.InfoRetriever), new(*application.InfoRetriever)),
	wire.Bind(new(apiProfile.Updater), new(*application.UpdateProfile)),
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
	session.Set,
	account.Set,
	ConfigSet,
	RouteSet,
	DatabaseSet,
)

var RouteSet = wire.NewSet(
	ProvideRouteResolver,
	wire.Bind(new(apiPacket.Resolver), new(*packet.Resolver)),
)

var ConfigSet = wire.NewSet(
	ProvideConfig,
	wire.Bind(new(apiConfig.Reader), new(*viper.Viper)),
)

var DatabaseSet = wire.NewSet(
	ProvideDatabase,
)

// singletons
var (
	routeResolver     *packet.Resolver
	routeResolverOnce sync.Once

	viperInstance *viper.Viper
	viperOnce     sync.Once

	databaseConnection     *sqlx.DB
	databaseConnectionOnce sync.Once
)

func ProvideRouteResolver() *packet.Resolver {
	routeResolverOnce.Do(func() {
		routeResolver = &packet.Resolver{}
	})

	return routeResolver
}

func ProvideConfig() *viper.Viper {
	viperOnce.Do(func() {
		v, err := config.Build(
			config.WithEnvFile(".env"),
			config.WithConfigDirectory("config/"))

		if err != nil {
			log.Fatal().Msgf("unable to initialize viper: %s", err.Error())
		}

		viperInstance = v
	})

	return viperInstance
}

func ProvideDatabase(config apiConfig.Reader) *sqlx.DB {
	databaseConnectionOnce.Do(func() {
		driver := config.GetString("database.driver")
		if driver != "mysql" {
			log.Fatal().Msgf("Database driver '%s' is unsupported, you can only use 'mysql' at this moment")
		}

		port, err := strconv.Atoi(config.GetString("database.drivers.mysql.port"))
		if err != nil {
			log.Fatal().Err(err)
		}

		conn, err := database.NewMySQLConnection(database.ConnectionInfo{
			Host:     config.GetString("database.drivers.mysql.host"),
			User:     config.GetString("database.drivers.mysql.user"),
			Password: config.GetString("database.drivers.mysql.password"),
			Port:     port,
			DBName:   config.GetString("database.drivers.mysql.dbname"),
		})

		if err != nil {
			log.Fatal().Msgf("Unable to connect to the database! Reason: %s", err.Error())
		}

		log.Info().Msg("Configured database pool")

		databaseConnection = conn
	})

	return databaseConnection
}

func ProvideNameCheckHandler(availableFunc availability.UsernameAvailableFunc) registration2.NameCheckHandler {
	return registration2.NewNameCheckHandler(availableFunc)
}

func NewApp(packetResolver apiPacket.Resolver, server *connection.GameSocket, viper *viper.Viper, db *sqlx.DB) *App {
	return &App{PacketResolver: packetResolver, GameServer: server, Config: viper, DB: db}
}

func NewNameCheckHandler(availableFunc availability.UsernameAvailableFunc) registration2.NameCheckHandler {
	return registration2.NewNameCheckHandler(availableFunc)
}

func NewPasswordCheckHandler(validationFunc password.ValidationFunc) registration2.PasswordVerifyHandler {
	return registration2.PasswordVerifyHandler{PasswordValidator: validationFunc}
}

func ProvideRegistrationService(createAccount apiAccount.CreateAccount, createProfile apiProfile.CreateProfile) saga.RegistrationService {
	return saga.RegistrationService{
		CreateAccount: createAccount,
		CreateProfile: createProfile,
	}
}

func NewRegisterHandler(registrationService saga.RegistrationService) register.Handler {
	return register.Handler{
		RegistrationService: registrationService,
	}
}

func ProvideValidateUsernameMiddleware(availableFunc availability.UsernameAvailableFunc) middleware.ValidateUsername {
	return middleware.ValidateUsername{
		AvailabilityChecker: availableFunc,
	}
}

func ProvideLoginService(credentialsVerifier apiAccount.VerifyCredentials, store *session.Store) saga.LoginService {
	return saga.LoginService{
		CredentialsVerifier: credentialsVerifier,
		SessionStore:        store,
	}
}

func ProvideTryLoginHandler(service saga.LoginService) login.TryLoginHandler {
	return login.TryLoginHandler{
		LoginService: service,
	}
}

func ProvideUserInfoHandler(store *session.Store, retriever apiProfile.InfoRetriever) user.InfoHandler {
	return user.InfoHandler{
		SessionStore:  store,
		InfoRetriever: retriever,
	}
}

func ProvideUpdateUserHandler(updater apiProfile.Updater, store *session.Store) user.Update {
	return user.Update{
		ProfileUpdater: updater,
		SessionStore:   store,
	}
}

func InitializeUpdateUserHandler() user.Update {
	wire.Build(ProvideUpdateUserHandler, ProfileSet, session.Set, ConfigSet, DatabaseSet)

	return user.Update{}
}

func InitializeValidateUsernameMiddleware() middleware.ValidateUsername {
	wire.Build(ProvideValidateUsernameMiddleware, account.Set, ConfigSet, DatabaseSet)

	return middleware.ValidateUsername{}
}

func InitializeRegisterHandler() register.Handler {
	wire.Build(NewRegisterHandler, ProvideRegistrationService, ProfileSet, account.Set, ConfigSet, DatabaseSet)

	return register.Handler{}
}

func InitializeNameCheckHandler() registration2.NameCheckHandler {
	wire.Build(NewNameCheckHandler, account.Set, ConfigSet, DatabaseSet)

	return registration2.NameCheckHandler{}
}

func InitializePasswordVerifyHandler() registration2.PasswordVerifyHandler {
	wire.Build(NewPasswordCheckHandler, account.Set)

	return registration2.PasswordVerifyHandler{}
}

func ProvideLoginAfterRegistrationMiddleware(service saga.LoginService) middleware.LoginAfterRegistration {
	return middleware.LoginAfterRegistration{
		LoginService: service,
	}
}

func InitializeLoginAfterRegistrationMiddleware() middleware.LoginAfterRegistration {
	wire.Build(ProvideLoginAfterRegistrationMiddleware, ProvideLoginService, session.Set, account.Set, ConfigSet, DatabaseSet)

	return middleware.LoginAfterRegistration{}
}

func InitializeTryLoginHandler() login.TryLoginHandler {
	wire.Build(ProvideTryLoginHandler, ProvideLoginService, session.Set, account.Set, ConfigSet, DatabaseSet)

	return login.TryLoginHandler{}
}

func InitializeInfoRetrieveHandler() user.InfoHandler {
	wire.Build(ProvideUserInfoHandler, session.Set, ProfileSet, ConfigSet, DatabaseSet)

	return user.InfoHandler{}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)
	return &App{}
}
