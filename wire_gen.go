// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"hc/api/account/availability"
	"hc/api/account/password"
	"hc/api/config"
	"hc/api/packet"
	"hc/internal/account"
	"hc/internal/connection"
	packet2 "hc/internal/packet"
	config2 "hc/pkg/config"
	"hc/pkg/database"
	"hc/presentationlayer/incoming/registration"
	"strconv"
	"sync"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func InitializeNameCheckHandler() registration.NameCheckHandler {
	viper := ProvideConfig()
	db := ProvideDatabase(viper)
	player := account.ProvidePlayerStore(db)
	usernameAvailableFunc := account.ProvideCheckNameAvailabilityHandler(player)
	nameCheckHandler := NewNameCheckHandler(usernameAvailableFunc)
	return nameCheckHandler
}

func InitializePasswordVerifyHandler() registration.PasswordVerifyHandler {
	validationFunc := account.ProvidePasswordValidator()
	passwordVerifyHandler := NewPasswordCheckHandler(validationFunc)
	return passwordVerifyHandler
}

func InitializeApp() *App {
	resolver := ProvideRouteResolver()
	wrapFunc := connection.ProvideMiddlewareWrapper()
	frontController := connection.ProvideFrontController(resolver, wrapFunc)
	repository := connection.ProvideSocketRepository(frontController)
	trafficParser := connection.ProvideTrafficParser()
	requestPool := connection.ProvideRequestPool()
	trafficManager := connection.ProvideTrafficManager(repository, trafficParser, requestPool)
	gameSocket := connection.ProvideGameSocket(repository, trafficManager)
	viper := ProvideConfig()
	db := ProvideDatabase(viper)
	app := NewApp(resolver, gameSocket, viper, db)
	return app
}

// wire.go:

var AppSet = wire.NewSet(connection.GameServerSet, account.Set, ConfigSet,
	RouteSet,
	DatabaseSet,
)

var RouteSet = wire.NewSet(
	ProvideRouteResolver, wire.Bind(new(packet.Resolver), new(*packet2.Resolver)),
)

var ConfigSet = wire.NewSet(
	ProvideConfig, wire.Bind(new(config.Reader), new(*viper.Viper)),
)

var DatabaseSet = wire.NewSet(
	ProvideDatabase,
)

// singletons
var (
	routeResolver     *packet2.Resolver
	routeResolverOnce sync.Once

	viperInstance *viper.Viper
	viperOnce     sync.Once

	databaseConnection     *sqlx.DB
	databaseConnectionOnce sync.Once
)

func ProvideRouteResolver() *packet2.Resolver {
	routeResolverOnce.Do(func() {
		routeResolver = &packet2.Resolver{}
	})

	return routeResolver
}

func ProvideConfig() *viper.Viper {
	viperOnce.Do(func() {
		v, err := config2.Build(config2.WithEnvFile(".env"), config2.WithConfigDirectory("config/"))

		if err != nil {
			log.Fatal().Msgf("unable to initialize viper: %s", err.Error())
		}

		viperInstance = v
	})

	return viperInstance
}

func ProvideDatabase(config3 config.Reader) *sqlx.DB {
	databaseConnectionOnce.Do(func() {
		driver := config3.GetString("database.driver")
		if driver != "mysql" {
			log.Fatal().Msgf("Database driver '%s' is unsupported, you can only use 'mysql' at this moment")
		}

		port, err := strconv.Atoi(config3.GetString("database.drivers.mysql.port"))
		if err != nil {
			log.Fatal().Err(err)
		}

		conn, err := database.NewMySQLConnection(database.ConnectionInfo{
			Host:     config3.GetString("database.drivers.mysql.host"),
			User:     config3.GetString("database.drivers.mysql.user"),
			Password: config3.GetString("database.drivers.mysql.password"),
			Port:     port,
			DBName:   config3.GetString("database.drivers.mysql.dbname"),
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

func NewApp(packetResolver packet.Resolver, server *connection.GameSocket, viper2 *viper.Viper, db *sqlx.DB) *App {
	return &App{PacketResolver: packetResolver, GameServer: server, Config: viper2, DB: db}
}

func NewNameCheckHandler(availableFunc availability.UsernameAvailableFunc) registration.NameCheckHandler {
	return registration.NewNameCheckHandler(availableFunc)
}

func NewPasswordCheckHandler(validationFunc password.ValidationFunc) registration.PasswordVerifyHandler {
	return registration.PasswordVerifyHandler{PasswordValidator: validationFunc}
}
