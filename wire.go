//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	apiConfig "hc/api/config"
	apiPacket "hc/api/packet"
	"hc/internal/connection"
	"hc/internal/packet"
	"hc/pkg/config"
	"hc/pkg/database"
	"strconv"
	"sync"
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
	ProvideRouteResolver,
	ProvideConfig,
	wire.Bind(new(apiPacket.Resolver), new(*packet.Resolver)),
	wire.Bind(new(apiConfig.Reader), new(*viper.Viper)),
	ProvideDatabase,
)

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
		routeResolver = packet.NewResolver(CollectRoutes())
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

		log.Info().Msg("Configured database pool")

		databaseConnection = conn
	})

	return databaseConnection
}

func NewApp(server *connection.GameSocket, viper *viper.Viper) *App {
	return &App{GameServer: server, Config: viper}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)
	return &App{}
}
