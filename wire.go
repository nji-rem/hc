//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	apiConfig "hc/api/config"
	apiPacket "hc/api/packet"
	"hc/internal/connection"
	"hc/internal/packet"
	"hc/pkg/config"
	"sync"
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
	ProvideRouteResolver,
	ProvideConfig,
	wire.Bind(new(apiPacket.Resolver), new(*packet.Resolver)),
	wire.Bind(new(apiConfig.Reader), new(*viper.Viper)),
)

var (
	routeResolver     *packet.Resolver
	routeResolverOnce sync.Once

	viperInstance *viper.Viper
	viperOnce     sync.Once
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

func NewApp(server *connection.GameSocket, viper *viper.Viper) *App {
	return &App{GameServer: server, Config: viper}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)
	return &App{}
}
