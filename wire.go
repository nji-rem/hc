//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"hc/api/account/availability"
	"hc/api/account/password"
	apiConfig "hc/api/config"
	apiPacket "hc/api/packet"
	"hc/internal/account"
	"hc/internal/connection"
	"hc/internal/packet"
	"hc/pkg/config"
	"hc/pkg/database"
	"hc/presentationlayer/incoming/registration"
	"hc/presentationlayer/incoming/registration/register"
	"strconv"
	"sync"
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
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

func ProvideNameCheckHandler(availableFunc availability.UsernameAvailableFunc) registration.NameCheckHandler {
	return registration.NewNameCheckHandler(availableFunc)
}

func NewApp(packetResolver apiPacket.Resolver, server *connection.GameSocket, viper *viper.Viper, db *sqlx.DB) *App {
	return &App{PacketResolver: packetResolver, GameServer: server, Config: viper, DB: db}
}

func NewNameCheckHandler(availableFunc availability.UsernameAvailableFunc) registration.NameCheckHandler {
	return registration.NewNameCheckHandler(availableFunc)
}

func NewPasswordCheckHandler(validationFunc password.ValidationFunc) registration.PasswordVerifyHandler {
	return registration.PasswordVerifyHandler{PasswordValidator: validationFunc}
}

func NewRegisterHandler() register.Handler {
	return register.Handler{}
}

func InitializeRegisterHandler() register.Handler {
	wire.Build(NewRegisterHandler)

	return register.Handler{}
}

func InitializeNameCheckHandler() registration.NameCheckHandler {
	wire.Build(NewNameCheckHandler, account.Set, ConfigSet, DatabaseSet)

	return registration.NameCheckHandler{}
}

func InitializePasswordVerifyHandler() registration.PasswordVerifyHandler {
	wire.Build(NewPasswordCheckHandler, account.Set)

	return registration.PasswordVerifyHandler{}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)
	return &App{}
}
